package database

import (
	"testing"
	"github.com/XMLPro/todoControl/models"
	"log"
)

//取り出したタプルが妥当か
func TestDB_GetAllTasks(t *testing.T) {
	db := NewDB("test.db")
	defer db.Close()
	tasks, err := db.GetAllTasks()

	if err != nil {
		t.Error(err)
	}

	for _, task := range tasks {
		if err := task.Valid(); err != nil {
			t.Error("妥当じゃないデータにマッピングされてる", err)
		}
	}
}

//新規挿入はtaskidが同じでも無視する
func TestDB_InsertTask_(t *testing.T) {
	db := NewDB("test.db")
	defer db.Close()

	testTask := models.Task{
		TaskId:     0,
		Content:    "Content1",
		PlaceId:    0,
		Importance: 0,
		Done:       false,
	}

	err := db.InsertTask(testTask)
	if err != nil {
		t.Error("正しく挿入されてない", err)
	}

	testTask2 := models.Task{
		TaskId:     0,
		Content:    "Content2",
		PlaceId:    0,
		Importance: 0,
		Done:       false,
	}

	err = db.InsertTask(testTask2)
	if err != nil {
		t.Error("taskIdが無視されてない",err)
	}
}

func TestDB_UpdateTask(t *testing.T){
	db := NewDB("test.db")
	defer db.Close()

	updateTask := models.Task{
		TaskId: 0,
		Content: "update target task",
		PlaceId: 0,
		Importance: 0,
		Done: true,
	}

	if err := db.InsertTask(updateTask); err != nil{
		t.Log(err)
	}
}


func TestDB_DeleteDoneTask(t *testing.T) {
	db := NewDB("test.db")
	defer db.Close()

	testTask := models.Task{
		TaskId: 0,
		Content: "delete target task",
		PlaceId: 0,
		Importance: 0,
		Done: true,
	}

	if err := db.InsertTask(testTask); err != nil{
		t.Log(err)
	}

	//delete point
	db.DeleteDoneTask()

	tasks, err  := db.GetAllTasks()
	if err != nil{
		t.Log(err)
	}

	for _, task := range tasks {
		t.Log(task)
		if task.Done {
			t.Error("Done: trueなタプルが消されてない", task)
		}
	}
}
