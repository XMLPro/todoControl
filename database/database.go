package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"log"
)

type DB struct {
	db *sql.DB
}

func NewDB(dbName string) *DB {
	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatal(err)
	}

	return &DB{db:db}
}

func (db *DB) Close() {
	db.db.Close()
}
