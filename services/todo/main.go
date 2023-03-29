package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"services/todo/db"
	. "services/todo/handler"
)

func main() {
	db.Init()

	router := mux.NewRouter()

	router.HandleFunc("/getTask/{id}", GetTask).Methods("GET")
	router.HandleFunc("/getTasks", GetTasks).Methods("GET")
	router.HandleFunc("/createTask", CreateTask).Methods("POST")
	router.HandleFunc("/", Health).Methods("GET")

	http.ListenAndServe("127.0.0.1:3000", router)
}
