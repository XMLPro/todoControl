package database

import (
	"github.com/XMLPro/todoControl/models"
	"log"
	sq "github.com/Masterminds/squirrel"
)

func (db *DB) GetTasks() []models.Task {
	rows, err := sq.Select("*").
		From("tasks").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Fatal(err)
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.Id,
			&task.Content,
			&task.LimitTime,
			&task.Workload,
			&task.PlaceId,
			&task.Importance,
		)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (db *DB) InsertTask(task models.Task) {
	_, err := sq.Insert("tasks").
		Columns("content", "limit_time", "workload", "importance").
		Values(
		task.Content,
		task.LimitTime,
		task.Workload,
		task.Importance,
		).
		RunWith(db.db).
		Exec()

	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) GetLatestTask() models.Task {
	row, err := sq.Select("*").
		From("tasks").
		Where("id=(select max(id) from tasks)").
		RunWith(db.db).
		Query()

	if err != nil {
		log.Fatal(err)
	}

	var task models.Task
	if row == nil {
		return task
	}

	row.Next()
	err = row.Scan(
		&task.Id,
		&task.Content,
		&task.LimitTime,
		&task.Workload,
		&task.PlaceId,
		&task.Importance,
	)

	if err != nil{
		log.Fatal(err)
	}

	return task
}
