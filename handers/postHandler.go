package handers

import (
	"net/http"
	"github.com/XMLPro/todoControl/database"
	"encoding/json"
	"log"
)


func GetTasksHandler(w http.ResponseWriter,  r *http.Request) {
	db  := database.NewDB("dev.db")

	encJson, err := json.Marshal(db.GetAllTasks())

	log.Println(string(encJson))

	if err != nil{
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encJson)
}