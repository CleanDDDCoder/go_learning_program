package validation

import (
	"errors"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{name: "valid", input: "gopher", expectedErr: nil},
		{name: "empty", input: "", expectedErr: ErrEmptyUsername},
		{name: "short", input: "go", expectedErr: ErrShortUsername},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.input)
			if tt.expectedErr == nil {
				if err != nil {
					t.Fatalf("ValidateUsername(%q) err = %v, want nil", tt.input, err)
				}
			} else {
				if err == nil {
					t.Fatalf("ValidateUsername(%q) err = nil, want %v", tt.input, tt.expectedErr)
				}
				if !errors.Is(err, tt.expectedErr) {
					t.Fatalf("ValidateUsername(%q) err = %v, want %v", tt.input, err, tt.expectedErr)
				}
			}
		})
	}
}