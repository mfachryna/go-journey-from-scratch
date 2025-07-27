package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow, Backend!")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow, Fachry!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hellow, ", name)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := map[string]string{"message": "Success!", "method": "GET"}
		json.NewEncoder(w).Encode(data)
	case http.MethodPost:
		data := map[string]string{"message": "Success!", "method": "POST"}
		json.NewEncoder(w).Encode(data)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/name", nameHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/api", apiHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
