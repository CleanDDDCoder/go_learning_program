// Package internalpackages demonstrates Go's internal package restriction
package internalpackages

import (
	"fmt"
	
	"go_learning_program/lessons/15-tooling-and-release/03-internal-packages/internal"
)

// GetInternalDemo demonstrates using an internal package
func GetInternalDemo() string {
	service := internal.NewInternalService("DemoService")
	return service.Process()
}

// GetServiceName returns the name from internal package
func GetServiceName() string {
	service := internal.NewInternalService("DemoService")
	return service.GetName()
}

func init() {
	fmt.Println("internalpackages package initialized")
}