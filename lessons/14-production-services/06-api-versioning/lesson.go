package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// UserV1 represents a user in API v1
type UserV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserV2 represents a user in API v2
type UserV2 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/v1/", v1Handler)
	http.HandleFunc("/v2/", v2Handler)
	http.ListenAndServe(":8080", nil)
}

func v1Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v1/")
	if path == "users" {
		users := []UserV1{
			{ID: 1, Name: "Alice"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func v2Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v2/")
	if path == "users" {
		users := []UserV2{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
}