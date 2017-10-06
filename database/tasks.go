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
