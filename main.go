package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path"
)

func doNothing(w http.ResponseWriter, r *http.Request) {}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("template", "question.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := OpenDb(path.Join("database", "words.db"))
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	word, err := GetRandWord(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ipv4 string = r.Header.Get("X-FORWARDED-FOR")
	if len(ipv4) < 1 {
		ipv4, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	fmt.Printf("> loading word %s (id:%d) on IP: %s\n", word.Str, word.ID, ipv4)
	if err := tmpl.Execute(w, word); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func initialize() {
	fmt.Println("> Parse data file")
	items, err := ParseData()

	if err != nil {
		panic(err)
	}

	fmt.Println("> Creating db")
	fp := path.Join("database", "words.db")
	os.MkdirAll("./database/", 0755)
	os.Remove(fp)
	os.Create(fp)

	db, err := OpenDb(fp)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("> Populate db")
	InitDb(db, items)
	fmt.Println("> Finished populating the db")
}

func main() {
	initialize()

	fmt.Println("> Http Server started")
	http.HandleFunc("/", getQuestion)
	http.HandleFunc("/favicon.ico", doNothing)
	http.ListenAndServe(":8080", nil)
}
