package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

var tasks []Task = make([]Task, 0)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func taskAddHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &task)

	fmt.Println(task)
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
