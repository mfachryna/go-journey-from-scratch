package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:id`
	Name string `json:name`
}

var users = []User{
	{ID: 1, Name: "Muhammad"},
	{ID: 2, Name: "Fachry"},
	{ID: 3, Name: "Noorchoolish"},
	{ID: 4, Name: "Arif"},
}

var nextID = 4

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		parts := strings.Split(r.URL.Path, "/")

		if len(parts) == 3 && parts[2] != "" {
			id, err := strconv.Atoi(parts[2])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
				return
			}

			for _, user := range users {
				if user.ID == id {
					json.NewEncoder(w).Encode(user)
					return
				}
			}

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		} else {
			json.NewEncoder(w).Encode(users)
		}
	case http.MethodPost:
		newUser := User{ID: nextID, Name: "New User"}
		users = append(users, newUser)
		nextID++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")

	}
}

func main() {
	http.HandleFunc("/users/", usersHandler)

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
