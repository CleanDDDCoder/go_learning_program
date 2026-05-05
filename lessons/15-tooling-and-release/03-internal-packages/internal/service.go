// Package internal demonstrates Go's internal package restriction
package internal

import (
	"fmt"
)

// InternalService represents a service with internal state
type InternalService struct {
	name string
}

// NewInternalService creates a new internal service
func NewInternalService(name string) *InternalService {
	return &InternalService{name: name}
}

// GetName returns the service name
func (s *InternalService) GetName() string {
	return s.name
}

// Process runs internal business logic
func (s *InternalService) Process() string {
	return fmt.Sprintf("Processing in %s", s.name)
}

func init() {
	fmt.Println("internal package initialized")
}