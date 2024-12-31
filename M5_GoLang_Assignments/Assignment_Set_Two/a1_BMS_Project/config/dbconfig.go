package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./myusers.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
