package database

import (
	"testing"
	"fmt"
	"github.com/XMLPro/todoControl/models"
)

func TestDB_GetAllTasks(t *testing.T) {
	db := NewDB("../dev.db")
	tasks := db.GetAllTasks()

	if tasks == nil {
		t.Error(tasks)
	}

	fmt.Println(tasks)
}

func TestDB_InsertTask(t *testing.T) {
	db := NewDB("../dev.db")
	task := models.Task{Task_id:0, Content:"Aaa", Place_id:0, Importance:0}

	err := db.InsertTask(task)

	if err != nil {
		t.Error(err)
	}
}
