package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// Widget is the domain object exposed by the starter service.
type Widget struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Store persists widgets behind a testable boundary.
type Store interface {
	Save(ctx context.Context, widget Widget) error
}

// Handler returns an HTTP handler for the service.
func Handler(store Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /widgets", func(w http.ResponseWriter, r *http.Request) {
		var widget Widget
		if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		// TODO: Validate fields, call Store, and preserve cancellation behavior.
		w.WriteHeader(http.StatusAccepted)
	})
	return mux
}
