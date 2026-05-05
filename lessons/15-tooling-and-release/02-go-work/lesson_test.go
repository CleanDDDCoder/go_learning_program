package goWork

import (
	"testing"
)

func TestGetWorkspaceInfo(t *testing.T) {
	info := GetWorkspaceInfo()
	if info == "" {
		t.Error("GetWorkspaceInfo returned empty string")
	}
}

func TestExampleFunction(t *testing.T) {
	result := ExampleFunction()
	if result == "" {
		t.Error("ExampleFunction returned empty string")
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"negative numbers", -2, -3, 6},
		{"mixed numbers", -2, 3, -6},
		{"zeros", 0, 5, 0},
		{"one", 1, 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}