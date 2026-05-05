package main

import (
	"errors"
	"fmt"
)

// Exercise: Error Wrapping
// Practice using fmt.Errorf with %w to wrap errors

var ErrNotFound = errors.New("resource not found")
var ErrUnauthorized = errors.New("unauthorized access")

// FindUser wraps ErrNotFound with context using %w
// TODO: Implement this function
func FindUser(id string) error {
	if id == "" {
		// Use fmt.Errorf with %w to wrap ErrNotFound with additional context
		return fmt.Errorf("invalid user id: %w", ErrNotFound)
	}
	// Simulate user not found case
	if id == "unknown" {
		return fmt.Errorf("user lookup failed for id %q: %w", id, ErrNotFound)
	}
	return nil
}

// GetUserProfile wraps the error from FindUser with more context
// TODO: Implement this function to wrap multiple errors
func GetUserProfile(id string) error {
	err := FindUser(id)
	if err != nil {
		// Wrap the existing error with additional context
		return fmt.Errorf("failed to get profile for user %q: %w", id, err)
	}
	return nil
}

// HandleUserRequest uses errors.Is to check for specific error types
// TODO: Implement this function to use errors.Is for error checking
func HandleUserRequest(id string) error {
	err := GetUserProfile(id)
	if err != nil {
		// Use errors.Is to check if the error chain contains ErrNotFound
		if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("cannot handle request: %w", err)
		}
		return err
	}
	fmt.Printf("Successfully handled request for user %s\n", id)
	return nil
}
