//go:build ignore

package validation

import "errors"

// Validation errors returned by ValidateUsername.
var (
	ErrEmptyUsername = errors.New("empty username")
	ErrShortUsername = errors.New("username too short")
)

// ValidateUsername checks the username rules.
func ValidateUsername(username string) error {
	if username == "" {
		return ErrEmptyUsername
	}
	if len(username) < 3 {
		return ErrShortUsername
	}
	return nil
}
