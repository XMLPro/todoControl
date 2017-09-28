package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

type DB struct {
	db *sql.DB
}

func NewDB(dbname string) *DB {
	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		panic(err)
	}

	return &DB{db: db}
}

func (db *DB) Close() {
	db.db.Close()
}


/*

func (db *DB) DeleteDoneTask() error {
	return nil
}


func (db *DB) existTaskRecord(id int){

}

func (db *DB) ExistiRecoed(table string, id int) bool {

	return true
}
*/
