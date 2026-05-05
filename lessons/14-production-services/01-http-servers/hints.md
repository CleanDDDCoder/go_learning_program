# HTTP Servers

Create a basic HTTP server that handles different routes.

## Objectives

- Understand Go's `net/http` package
- Create handlers for different routes
- Return appropriate HTTP status codes

## Setup

The student should create a file `server.go` with a basic HTTP server.

## Exercise

Create an HTTP server that:

1. Handles GET requests to `/hello` returning "Hello, World!"
2. Handles GET requests to `/health` returning 200 OK with `{"status": "ok"}`
3. Returns 404 for any unknown route

## Hints

- Use `http.HandleFunc` to register handlers
- Use `http.ListenAndServe` to start the server
- Set the `Content-Type` header for JSON responses

## Solution

```go
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Hello, World!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```