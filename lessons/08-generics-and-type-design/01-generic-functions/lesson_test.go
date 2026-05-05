package main

import (
	"testing"
)

func TestSumInts(t *testing.T) {
	result := Sum([]int{1, 2, 3, 4, 5})
	if result != 15 {
		t.Errorf("Sum([]int{1,2,3,4,5}) = %d; want 15", result)
	}
}

func TestSumFloats(t *testing.T) {
	result := Sum([]float64{1.5, 2.5, 3.0})
	if result != 7.0 {
		t.Errorf("Sum([]float64{1.5,2.5,3.0}) = %v; want 7.0", result)
	}
}

func TestDoubleInt(t *testing.T) {
	result := Double(5)
	if result != 10 {
		t.Errorf("Double(5) = %v; want 10", result)
	}
}

func TestDoubleFloat(t *testing.T) {
	result := Double(3.5)
	if result != 7.0 {
		t.Errorf("Double(3.5) = %v; want 7.0", result)
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for _, tc := range tests {
		result := FizzBuzz(tc.input)
		if result != tc.expected {
			t.Errorf("FizzBuzz(%d) = %q; want %q", tc.input, result, tc.expected)
		}
	}
}
