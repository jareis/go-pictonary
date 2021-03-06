package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func loadDb() (*sql.DB, error) {
	fp := path.Join("database", "dictionary")

	fmt.Println("> Loading database at:" + fp)
	db, err := sql.Open("sqlite3", fp)
	if err != nil {
		return db, err
	}
	return db, err
}

func iniDB() {
	fp := path.Join("database", "dic.txt")
	file, err := os.Open(fp)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var line string

	db, err := loadDb()
	defer db.Close()

	_, err = db.Exec("DELETE FROM words")
	if err != nil {
		panic(err)
	}

	var i int
	i = 1
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) < 1 {
			fmt.Println("> Nothing more to do")
			break
		}

		// Process the line here.
		fmt.Printf("> Read %d characters\n", len(line))
		fmt.Printf(">> %s\n", line)
		_, err = db.Exec("INSERT INTO words (id, str, weight) VALUES ($1, $2, $3)", i, line, 0)
		i++

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		panic(err)
	}

	return
}

func queryDb() (Word, error) {
	db, err := loadDb()
	defer db.Close()
	w := Word{}

	if err != nil {
		return w, err
	}

	fmt.Println("> Select word")
	db.QueryRow("SELECT * FROM words ORDER BY RANDOM(), weight DESC LIMIT 1").Scan(&w.ID, &w.Str, &w.Weight)

	fmt.Printf("> Add weight to word(id=%d) \n", w.ID)
	_, err = db.Exec("UPDATE words SET weight=weight+1 WHERE id=$1", w.ID)
	if err != nil {
		return w, err
	}

	return w, err
}

func showQuestion(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("template", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	word, err := queryDb()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, word); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func main() {
	iniDB()
	fmt.Println("> Http Server started")
	http.HandleFunc("/", showQuestion)
	http.HandleFunc("/favicon.ico", doNothing)
	http.ListenAndServe(":8080", nil)
}
