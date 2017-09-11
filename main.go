package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"github.com/XMLPro/todoControl/database"
)

var db *database.DB = database.NewDB("test.db")

func main() {
	r := initRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", TemporaryHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	return router
}

type User struct {
	Name string
	Pass string
}

//database
func Authentication(name string) () {

}
// database experiment code
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

}

func TemporaryHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	log.Println(user)
}
