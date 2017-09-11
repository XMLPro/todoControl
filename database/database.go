package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DB struct {
	db *sql.DB
}

type User struct {
	Id       int
	Name     string
	Password string
}

func NewDB(dbname string) *DB {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}
	return &DB{db: db}
}

func (db *DB) GetUserPass(username string) string {
	s, err := db.db.Query("SELECT * FROM user WHERE name == ?", username)
	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	user := &User{}
	s.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
	)

	return user.Password
}

func (db *DB) Close() {
	db.db.Close()
}
