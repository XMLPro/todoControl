package models

import "errors"

type Task struct {
	TaskId     int    `json:"task_id"`
	Content    string `json:"content"`
	PlaceId    int    `json:"place_id"`
	Importance int    `json:"importance"`
	Done       bool   `json:"done"`
}

func (m *Task) Valid() error {
	if m.Content == ""{
		return errors.New("Task Error: content is none ")
	}

	if m.Importance < 0 || m.Importance > 100 {
		return errors.New("Task Error: importance is out of range ")
	}

	return nil
}
