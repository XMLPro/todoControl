package database

import (
	"log"
	"github.com/XMLPro/todoControl/models"
	sq "github.com/Masterminds/squirrel"
)

func (db *DB) GetAllTasks() ([]models.Task, error) {
	rows, err := sq.Select("*").
		From("tasks").
		RunWith(db.db).
		Query()

	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(
			&task.TaskId,
			&task.Content,
			&task.PlaceId,
			&task.Importance,
			&task.Done,
		)

		if err != nil {
			log.Println(err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (db *DB) InsertTask(m models.Task) error {
	_, err := sq.Insert("tasks").
		Columns("content", "place_id", "important").
		Values(m.Content, m.PlaceId, m.TaskId).
		RunWith(db.db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteDoneTask() {
}

func (db *DB) NextTaskID() int {
	res, err := sq.Select("max(task_id)").
		From("tasks").
		Query()

	if err != nil {
		log.Println(err)
	}

	var next int
	err = res.Scan(&next)
	if err != nil {
		log.Println(err)
	}

	return next + 1
}
