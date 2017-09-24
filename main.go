package main

import (
	"github.com/XMLPro/todoControl/handers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := initRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", handers.NewTemplateHandler("index.html")).Methods("GET")
	router.HandleFunc("/register", TemporaryHandler).Methods("POST")
	router.HandleFunc("/task/all", handers.GetTasksHandler).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return router
}

func TemporaryHandler(w http.ResponseWriter, r *http.Request) {
}
