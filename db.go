package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDb Opens the DB to be used
//
func OpenDb(fp string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fp)
	if err != nil {
		return db, err
	}

	if db == nil {
		panic("DB is nil")
	}

	return db, err
}

// InitDb Initiates the DB to a usable state
//
func InitDb(db *sql.DB, items []string) {
	sql := `
    DROP TABLE IF EXISTS words;
    CREATE TABLE words (
        id integer,
        str varchar UNIQUE,
        weight integer NOT NULL DEFAULT 0,
        PRIMARY KEY (id)
    );
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	stmt, _ := db.Prepare("INSERT or IGNORE INTO words (str) VALUES (?);")
	defer stmt.Close()

	for _, item := range items {
		_, err = stmt.Exec(item)
		if err != nil {
			panic(err)
		}
	}
}

// GetRandWord get a random item
func GetRandWord(db *sql.DB) (Word, error) {
	w := Word{}
	sql := `
    SELECT id, str, weight
    FROM words
    ORDER BY RANDOM(), weight DESC LIMIT 1
    `
	err := db.QueryRow(sql).Scan(&w.ID, &w.Str, &w.Weight)
	db.Exec("UPDATE words SET weight=weight+1 WHERE id=$1", w.ID)

	return w, err
}
