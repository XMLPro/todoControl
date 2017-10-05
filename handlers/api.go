package handlers

import (
	"net/http"

	"strconv"
	"encoding/json"
)

type Task struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	Limit      string `json:"limit"`
	PlaceId    int    `json:"place_id"`
	Importance int    `json:"importance"`
}

var tasks []Task = initTask()

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	writeJson(w, jsonData)
}

func initTask() []Task {
	var tasks []Task

	for i := 1; i < 5; i++ {
		tasks = append(tasks, Task{Id: i, Content: "content" + strconv.Itoa(i), Limit: "-", PlaceId: 0, Importance: 0})
	}

	return tasks
}

func writeJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
