package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	. "services/todo/models"
	"services/todo/repositories"
	"strconv"
)

// GetTask Handler for returning a single task
func GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming Request on /getTask/:id")

	vars := mux.Vars(r)
	urlId, success := vars["id"]

	if !success {
		fmt.Println("Could not get id from request")
	}

	id, err := strconv.Atoi(urlId)

	if err != nil {
		fmt.Println("Could not convert id to int")
	}

	w.Header().Set("Content-Type", "application/json")

	task, err := repositories.GetTaskById(id)
	if err != nil {
		fmt.Println("Could not get task from database")
	}

	json.NewEncoder(w).Encode(task)
}

// GetTasks Handler for returning a slice of tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming Request on /getTasks")

	w.Header().Set("Content-Type", "application/json")

	tasks, err := repositories.GetAllTasks()
	if err != nil {
		fmt.Println("Could not get tasks from database")
	}

	json.NewEncoder(w).Encode(tasks)
}

// CreateTask Handler for creating a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming Request on /createTask")

	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	newTask, err := repositories.CreateTask(task)
	if err != nil {
		fmt.Println("Could not create task in database")
	}

	json.NewEncoder(w).Encode(newTask)
}
