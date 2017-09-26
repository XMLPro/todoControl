package database

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"

	"github.com/XMLPro/todoControl/models"
	"log"
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


func (db *DB) GetAllTasks() []models.Task {
	rows, err := sq.Select("*").
		From("tasks").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Println(err)
	}

	var tasks []models.Task

	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(
			&task.Task_id,
			&task.Content,
			&task.Place_id,
			&task.Importance,
		)

		if err != nil {
			log.Println(err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func (db *DB) GetAllPlaces() []models.Place {
	rows, err := sq.Select("*").
		From("places").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Println(err)
	}

	var places []models.Place

	for rows.Next() {
		place := models.Place{}
		err = rows.Scan(
			&place.Place_id,
			&place.Place_name,
			&place.Importance,
		)

		if err != nil {
			log.Println(err)
		}
		places = append(places, place)
	}

	return places
}

