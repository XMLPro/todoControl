package handers

import (
	"encoding/json"
	"net/http"

	"github.com/XMLPro/todoControl/database"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := database.NewDB("dev.db")

	encJson, err := json.Marshal(db.GetAllTasks())

	//log.Println(string(encJson))

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encJson)
}
