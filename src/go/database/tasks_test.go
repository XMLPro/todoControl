package database

import (
	"testing"
	"github.com/XMLPro/todoControl/src/goc/go/models"
	"log"
)

func TestDB_InsertTask(t *testing.T) {
	task := models.Task{
		Content:"a",
	}
	db := NewDB("test.db")
	db.InsertTask(task)

}

func TestDB_GetLatestTask(t *testing.T) {
	db := NewDB("test.db")
	log.Println(db.GetLatestTask())
}
