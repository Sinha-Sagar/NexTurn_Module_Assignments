package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	return db
}
