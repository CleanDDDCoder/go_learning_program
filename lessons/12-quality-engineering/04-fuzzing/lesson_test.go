package fuzzing

import (
	"testing"
)

// TestProcessFuzzTarget is a fuzz-style test that can be run with:
// go test -fuzz=FuzzProcess -fuzztime=10s
func TestProcessBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantErr  bool
		wantData string
	}{
		{"empty", []byte{}, false, ""},
		{"hello", []byte("hello"), false, "hello"},
		{"empty string", []byte(""), false, ""},
		{"numbers", []byte("123"), false, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Process(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.wantData {
				t.Errorf("Process(%v) = %v, want %v", tt.input, got, tt.wantData)
			}
		})
	}
}

// TestValidateEmailFuzzTarget tests email validation with generated inputs.
func TestValidateEmailProperties(t *testing.T) {
	// Property: valid emails should always return true
	// Property: invalid emails should return false
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"simple valid", "test@example.com", true},
		{"valid with dot", "john.doe@example.com", true},
		{"missing @", "testexample.com", false},
		{"empty", "", false},
		{"just @", "@", false},
		{"no domain", "test@", false},
		{"starts with @", "@test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateEmail(tt.email)
			if got != tt.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}
