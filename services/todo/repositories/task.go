package repositories

import (
	"log"
	. "services/todo/db"
	. "services/todo/models"
)

func CreateTask(task Task) (Task, error) {
	var newTask Task

	rows, err := Db.Query("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&newTask.ID, &newTask.Title, &newTask.Description, &newTask.Completed)
		if err != nil {
			log.Fatal(err)
		}
	}

	return newTask, nil
}

func GetAllTasks() (Tasks, error) {
	var tasks Tasks

	rows, err := Db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskById(id int) (Task, error) {
	var task Task

	rows, err := Db.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
	}

	return task, nil
}
