package types

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"positive", 42, "42"},
		{"zero", 0, "0"},
		{"negative", -123, "-123"},
		{"large positive", 1000000, "1000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToString(tt.input)
			if got != tt.want {
				t.Errorf("IntToString(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"valid positive", "42", 42},
		{"valid zero", "0", 0},
		{"valid negative", "-123", -123},
		{"invalid empty", "", 0},
		{"invalid text", "hello", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToInt(tt.input)
			if got != tt.want {
				t.Errorf("StringToInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}