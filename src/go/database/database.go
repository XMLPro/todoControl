package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB(db_name string) *DB {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db: db}
}
