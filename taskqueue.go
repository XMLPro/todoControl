package main

import (
	//"encoding/json"
	//"bytes"
	"github.com/gorilla/mux"
	"net/http"
)

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
