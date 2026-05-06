package main

import (
	"context"
	"fmt"
	"time"
)

// contextKey is a private type to avoid collisions with context values
type contextKey string

const requestIDKey contextKey = "requestID"

// TODO: Implement FetchData that accepts a context and respects cancellation
// If the context is cancelled, return early with context.Canceled error
func FetchData(ctx context.Context, id string) (string, error) {
	// Simulate work - check for context cancellation periodically
	select {
	case <-time.After(100 * time.Millisecond):
		return fmt.Sprintf("data for %s", id), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// TODO: Implement AddRequestID that creates a new context with a request ID value
// Use context.WithValue to add a key-value pair to the context
func AddRequestID(ctx context.Context, requestID string) context.Context {
	// Use context.WithValue to add the request ID
	return context.WithValue(ctx, requestIDKey, requestID)
}

// TODO: Implement GetRequestID that retrieves a request ID from a context
// Use ctx.Value to retrieve the request ID
func GetRequestID(ctx context.Context) string {
	// Retrieve the request ID from context
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return ""
}

func main() {
	// Test context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	result, err := FetchData(ctx, "test-id")
	fmt.Printf("Result: %s, Error: %v\n", result, err)

	// Test context with values
	ctx = AddRequestID(context.Background(), "req-123")
	reqID := GetRequestID(ctx)
	fmt.Printf("Request ID: %s\n", reqID)
}