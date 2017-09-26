package models

type Task struct {
	Task_id    int    `json:"task_id"`
	Content    string `json:"content"`
	Place_id   int    `json:"place_id"`
	Importance int    `json:"importance"`
}

func (m *Task) Valid() error {
	return nil
}

