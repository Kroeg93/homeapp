package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	. "services/todo/models"
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

	var dummyTask = Task{
		ID:          id,
		Title:       "Task 1",
		Description: "This is a task",
		Completed:   false,
	}

	json.NewEncoder(w).Encode(dummyTask)
}

// GetTasks Handler for returning a slice of tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming Request on /getTasks")

	var dummyTasks = []Task{
		{
			ID:          1,
			Title:       "Task 1",
			Description: "This is a task",
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Task 2",
			Description: "This is a task",
			Completed:   false,
		},
		{
			ID:          3,
			Title:       "Task 3",
			Description: "This is a task",
			Completed:   false,
		},
		{
			ID:          4,
			Title:       "Task 4",
			Description: "This is a task",
			Completed:   false,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummyTasks)
}
