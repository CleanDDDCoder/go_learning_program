# API Versioning

Design APIs that support versioning for backward compatibility.

## Objectives

- Implement URL-based API versioning
- Handle version-specific request handling
- Maintain backward compatibility

## Exercise

Create an API server that:

1. Supports `/v1/users` and `/v2/users` endpoints
2. V1 returns user data with `id`, `name` fields
3. V2 returns user data with `id`, `name`, `email` fields
4. Returns 404 for unsupported versions

## Hints

- Use a router that dispatches based on path prefix
- Create separate handler functions for each version
- Consider using a version header as alternative approach

## Solution

```go
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type UserV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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