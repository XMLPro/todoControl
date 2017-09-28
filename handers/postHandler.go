package handers

import (
	"encoding/json"
	"net/http"

	"github.com/XMLPro/todoControl/database"
	"log"
	"crypto/ecdsa"
)

const db_name  = "dev.db"

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	db := database.NewDB(db_name)

	encJson, err := json.Marshal(db.GetAllTasks())

	//log.Println(string(encJson))

	if err != nil {
		log.Fatal(err)
	}

	writeJson(w, encJson)
}

func GetPlaceHandler(w http.ResponseWriter, r *http.Request){
	db := database.NewDB(db_name)

	encJson, err := json.Marshal(db.GetAllPlaces())

	if err != nil {
		log.Fatal(err)
	}

	writeJson(w, encJson)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {



}

func writeJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
