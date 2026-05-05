package slices

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "empty", values: nil, want: 0},
		{name: "positive", values: []int{1, 2, 3}, want: 6},
		{name: "mixed", values: []int{-2, 5, -1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.values)
			if got != tt.want {
				t.Fatalf("Sum(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	got, ok := Last([]string{"go", "test"})
	if !ok || got != "test" {
		t.Fatalf("Last(non-empty) = %q, %v; want test, true", got, ok)
	}

	got, ok = Last(nil)
	if ok || got != "" {
		t.Fatalf("Last(nil) = %q, %v; want empty, false", got, ok)
	}
}
