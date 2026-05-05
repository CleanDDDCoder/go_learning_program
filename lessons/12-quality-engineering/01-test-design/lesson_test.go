//go:build ignore

package testdesign

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		want     int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"mixed numbers", -1, 2, 1},
		{"zero add", 0, 5, 5},
		{"both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.a, tt.b); got != tt.want {
				t.Errorf("Sum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		want     int
	}{
		{"a is greater", 5, 3, 5},
		{"b is greater", 2, 8, 8},
		{"equal values", 4, 4, 4},
		{"negative a greater", -1, -5, -1},
		{"negative b greater", -10, -2, -2},
		{"zero vs positive", 0, 5, 5},
		{"positive vs zero", 7, 0, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.a, tt.b); got != tt.want {
				t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		want     int
	}{
		{"a is smaller", 3, 5, 3},
		{"b is smaller", 8, 2, 2},
		{"equal values", 4, 4, 4},
		{"negative a smaller", -1, -5, -5},
		{"negative b smaller", -10, -2, -10},
		{"zero vs positive", 0, 5, 0},
		{"positive vs zero", 7, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.a, tt.b); got != tt.want {
				t.Errorf("Min(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		want     []string
	}{
		{"three elements", []string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{"single element", []string{"x"}, []string{"x"}},
		{"empty slice", []string{}, []string{}},
		{"two elements", []string{"a", "b"}, []string{"b", "a"}},
		{"already reversed", []string{"c", "b", "a"}, []string{"a", "b", "c"}},
		{"with duplicates", []string{"a", "a", "a"}, []string{"a", "a", "a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("Reverse(%v) length = %d, want %d", tt.input, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Reverse(%v)[%d] = %q, want %q", tt.input, i, got[i], tt.want[i])
				}
			}
		})
	}
}