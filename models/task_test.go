package models

import (
	"testing"
)

func TestTask_Valid(t *testing.T) {
	task := Task{TaskId: 0, Content: "content1", PlaceId: 0, Importance: 0, Done: false}
	if err := task.Valid(); err != nil{
		t.Error(err)
	}

	//Importance Error (0~100)
	task = Task{TaskId: 1, Content: "content2", PlaceId:-4, Importance: 400, Done: false}
	if err := task.Valid(); err == nil {
		t.Error(err)
	}

	//Content Error (no value)
	task = Task{TaskId: 3, Content: "", PlaceId: 0, Importance: 100, Done: false}
	if err := task.Valid(); err == nil {
		t.Error(err)
	}
}
