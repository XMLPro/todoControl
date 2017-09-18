package database

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
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

// これ分離したい
type Task struct {
	Task_id    int
	Content    string
	Place_id   int
	Importance int
}

type Place struct {
	Place_id   int
	Place_name string
	Importance string
}

func (db *DB) GetAllTasks() []Task {
	rows, err := sq.Select("*").
		From("tasks").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Println(err)
	}

	var tasks []Task

	for rows.Next() {
		task := Task{}
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

func (db *DB) GetAllPlaces() []Place {
	rows, err := sq.Select("*").
		From("places").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Println(err)
	}

	var places []Place

	for rows.Next() {
		place := Place{}
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

/*
func (db *DB) GetTask(task_id int) *Task {
	row, err := sq.Select("*").
		From("tasks").
		Where("task_id == ?", task_id).
		RunWith(db.db).
		Query()

	if err != nil {
		log.Println(err)
	}

	defer row.Close()

	var task *Task
	err = row.Scan(
		&task.Task_id,
		&task.Content,
		&task.Place_id,
		&task.Importance,
	)

	if err != nil {
		log.Println(err)
	}

	return task
}

func (db *DB) GetPlace(place_id int) (*Task, error) {
	row, err := sq.Select("*").
		From("tasks").
		Where("task_id == ?").
		Query()

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var task *Task
	err = row.Scan(
		&task.Task_id,
		&task.Content,
		&task.Place_id,
		&task.Importance,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}
*/
