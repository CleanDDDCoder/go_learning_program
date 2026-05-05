package goMod

import (
	"testing"
)

func TestGetModuleInfo(t *testing.T) {
	info := GetModuleInfo()
	if info == "" {
		t.Error("GetModuleInfo returned empty string")
	}
}

func TestExampleFunction(t *testing.T) {
	result := ExampleFunction()
	if result == "" {
		t.Error("ExampleFunction returned empty string")
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 2, 1},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}