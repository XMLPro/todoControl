package database

import (
	"testing"
	"fmt"
)

func TestDB_GetAllTasks(t *testing.T) {
	db := NewDB(test_db_name)
	tasks,err  := db.GetAllTasks()
	
	if err != nil{
		t.Error(err)
	}

	for task := range tasks {
		fmt.Println(task)
	}
}