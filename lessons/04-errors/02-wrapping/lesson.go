package wrapping

import "fmt"

// WrapReadError adds file context while preserving the original error.
func WrapReadError(filename string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("read %s: %w", filename, err)
}
