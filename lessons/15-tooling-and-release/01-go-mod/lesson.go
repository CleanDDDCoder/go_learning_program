// Package goMod demonstrates Go module dependency management
package goMod

import (
	"fmt"
)

// GetModuleInfo returns information about the current module
func GetModuleInfo() string {
	return "Go Modules and Dependency Management Lesson"
}

// ExampleFunction demonstrates basic module usage
func ExampleFunction() string {
	return "This is a demo of module dependency management"
}

// Sum adds two integers
func Sum(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("goMod package initialized")
}