package main

import (
	"errors"
	"fmt"
)

// TODO: Implement FetchUser that wraps the underlying error with context
// Return an error that wraps the original error using %w verb
func FetchUser(userID int) error {
	// Simulate an error
	return errors.New("database connection failed")
}

// TODO: Implement IsTemporaryError that uses errors.Is to check for temporary errors
func IsTemporaryError(err error) bool {
	// Use errors.Is to check if the error or any wrapped error matches a sentinel error
	return false
}

// TODO: Implement GetErrorCode that uses errors.As to extract a custom error type
type ErrorCode struct {
	Code    int
	Message string
}

func GetErrorCode(err error) (int, error) {
	// Use errors.As to extract the ErrorCode from the error chain
	return 0, errors.New("not implemented")
}

func main() {
	// Test error wrapping
	err := FetchUser(123)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Test error inspection
	if errors.Is(err, errors.New("database connection failed")) {
		fmt.Println("Found database error")
	}
}