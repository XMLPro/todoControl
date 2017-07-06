package main

import (
	//"encoding/json"
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	primaryKey int
	uname      string //uniqueな名前
	upass      string //hash
}

// json parse用構造体
type Task struct {
	//primaryKey int
	content    string `json:"content"`
	deadline   string `json:"deadline"`
	from       string `json:"from"`
	workload   string `json:"workload"`
	importance string `json:"importance"`
}

/*
	route
		/			- return index
		/task 		- get all task
			/add	- task add
			/delete - task delete
			/change - task change

		/user
			/signup - sign up
			/signin - sign in
			/signout - sign out
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)

	taskGroup := router.PathPrefix("/task").Subrouter()
	taskGroup.HandleFunc("/", taskAllHandler).Methods("POST")
	taskGroup.HandleFunc("/add", taskAddHandler).Methods("POST")
	taskGroup.HandleFunc("/delete", taskDeleteHandler).Methods("POST")
	taskGroup.HandleFunc("/change", taskChangeHandler).Methods("POST")

	/**
	userGroup := router.PathPrefix("/user").Subrouter()
	*/

	http.ListenAndServe(":8000", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func taskAddHandler(w http.ResponseWriter, r *http.Request) {
	responseWrite(w, r)
}

func taskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	responseWrite(w, r)
}

func taskChangeHandler(w http.ResponseWriter, r *http.Request) {
	responseWrite(w, r)
}

func taskAllHandler(w http.ResponseWriter, r *http.Request) {
	responseWrite(w, r)
}

func responseWrite(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.Write([]byte("non value"))
	}

	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	w.Write([]byte(bufbody.String()))
}
