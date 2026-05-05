package variables

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Alice", "Hello, Alice!"},
		{"Bob", "Hello, Bob!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.name)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
