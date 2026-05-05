// Package goWork demonstrates Go workspace management
package goWork

import (
	"fmt"
)

// GetWorkspaceInfo returns information about the current workspace
func GetWorkspaceInfo() string {
	return "Go Workspaces and Multi-Module Development"
}

// ExampleFunction demonstrates workspace usage
func ExampleFunction() string {
	return "This is a demo of workspace management"
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

func init() {
	fmt.Println("goWork package initialized")
}