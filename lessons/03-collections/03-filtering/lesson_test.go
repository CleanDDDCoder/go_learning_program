package filtering

import (
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{name: "mixed", values: []int{1, 2, 3, 4}, want: []int{2, 4}},
		{name: "none", values: []int{1, 3}, want: nil},
		{name: "negative", values: []int{-2, -1, 0}, want: []int{-2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvenNumbers(tt.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("EvenNumbers(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}
