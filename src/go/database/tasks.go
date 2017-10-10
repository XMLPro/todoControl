package database

import (
	"github.com/XMLPro/todoControl/src/go/models"
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

	if err != nil {
		log.Fatal(err)
	}

	return task
}

func (db *DB) UpdateTask(task models.Task){
	_, err := sq.Update("tasks").
		Set("content", task.Content).
		Set("limit_time", task.LimitTime).
		Set("workload", task.Workload).
		Set("place_id", task.PlaceId).
		Set("importance", task.Importance).
		Where("id=?", task.Id).
		RunWith(db.db).
		Exec()

	if err != nil {
		log.Fatal(err)
	}
}
