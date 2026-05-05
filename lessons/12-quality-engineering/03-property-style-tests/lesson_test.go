package propertytests

import (
	"reflect"
	"testing"
)

// TestReverseReverseProperties verifies that reversing twice returns the original slice.
// This is a property that should always hold regardless of input.
func TestReverseReverseProperties(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"empty", []int{}},
		{"single", []int{1}},
		{"two", []int{1, 2}},
		{"three", []int{1, 2, 3}},
		{"many", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Property: reversing twice should return the original
			got := ReverseReverse(ReverseReverse(tt.input))
			want := tt.input
			if !reflect.DeepEqual(got, want) {
				t.Errorf("ReverseReverse(ReverseReverse(%v)) = %v, want %v", tt.input, got, want)
			}
		})
	}
}

// TestSumIntsProperties verifies the sum property: sum of [1,2,3] + sum of [4,5,6]
// should equal sum of [1,2,3,4,5,6].
func TestSumIntsProperties(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	combined := []int{1, 2, 3, 4, 5, 6}

	got := SumInts(a) + SumInts(b)
	want := SumInts(combined)

	if got != want {
		t.Errorf("Sum property failed: %d + %d = %d, want %d", SumInts(a), SumInts(b), got, want)
	}
}