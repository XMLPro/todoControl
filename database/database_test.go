package database

import (
	"testing"
	"fmt"
)

func TestDB_GetAllTasks(t *testing.T) {
	db := NewDB("../dev.db")
	tasks := db.GetAllTasks()

	if tasks == nil {
		t.Error(tasks)
	}

	fmt.Println(tasks)
}
