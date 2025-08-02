package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{ID: 1, Title: "Learn Go basics", Completed: true},
	{ID: 2, Title: "Build a web server", Completed: true},
	{ID: 3, Title: "Create a REST API", Completed: false},
}
var nextID = 4

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 3 && parts[2] != "" {
			id, err := strconv.Atoi(parts[2])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
				return
			}
			for _, task := range tasks {
				if task.ID == id {
					json.NewEncoder(w).Encode(task)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		} else {
			json.NewEncoder(w).Encode(tasks)
		}

	case http.MethodPost:
		var newTask Task

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
			return
		}

		newTask.ID = nextID
		nextID++
		tasks = append(tasks, newTask)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	http.HandleFunc("/tasks/", tasksHandler)

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
