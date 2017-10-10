package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/XMLPro/todoControl/src/go/handlers"
)

func main() {
	r := initRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", handlers.NewTemplateHandler("index.html")).Methods("GET")
	r.HandleFunc("/task/all", handlers.GetTasksHandler).Methods("POST")
	r.HandleFunc("/task/add", handlers.AddTask).Methods("POST")
	r.HandleFunc("/task/update", handlers.UpdateTaskHandler).Methods("POST")
	r.HandleFunc("/place/all", handlers.GetPlacesHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}


