package benchmarking

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		wantSum  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"two elements", []int{1, 2}, 3},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, 1, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.numbers); got != tt.wantSum {
				t.Errorf("Sum(%v) = %v, want %v", tt.numbers, got, tt.wantSum)
			}
		})
	}
}