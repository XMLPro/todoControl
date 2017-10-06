package handlers

import (
	"net/http"

	"encoding/json"
	"github.com/XMLPro/todoControl/database"
	"log"
)

var db *database.DB = database.NewDB("dev.db")

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetTasks()
	jsonData, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	writeJson(w, jsonData)
}

func GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	places := db.GetPlaces()
	jsonData, err := json.Marshal(places)

	if err != nil {
		log.Fatal(err)
	}

	writeJson(w, jsonData)
}

func writeJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
