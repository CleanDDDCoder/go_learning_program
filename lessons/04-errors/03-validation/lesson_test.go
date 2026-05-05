package validation

import (
	"errors"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{name: "valid", input: "gopher", wantErr: nil},
		{name: "empty", input: "", wantErr: ErrEmptyUsername},
		{name: "short", input: "go", wantErr: ErrShortUsername},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ValidateUsername(%q) err = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
