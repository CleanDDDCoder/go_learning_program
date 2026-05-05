package main

import (
	"fmt"
	"html"
	"strings"
)

// Implement a function that validates and sanitizes user input.
// The function should:
// 1. Reject empty strings
// 2. Trim whitespace
// 3. Limit the length to 1000 characters
// 4. Escape HTML special characters to prevent XSS

func ValidateAndSanitize(input string) (string, error) {
	// Trim whitespace first
	trimmed := strings.TrimSpace(input)
	
	// Reject empty strings (including whitespace-only)
	if trimmed == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	
	// Limit the length to 1000 characters (after trimming)
	if len(trimmed) > 1000 {
		return "", fmt.Errorf("input exceeds maximum length")
	}
	
	// Escape HTML special characters to prevent XSS
	sanitized := html.EscapeString(trimmed)
	
	return sanitized, nil
}

func main() {
	fmt.Println("Hello, World!")
}