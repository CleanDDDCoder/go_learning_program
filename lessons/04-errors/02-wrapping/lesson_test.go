package wrapping

import (
	"errors"
	"strings"
	"testing"
)

func TestWrapReadError(t *testing.T) {
	base := errors.New("permission denied")
	err := WrapReadError("config.json", base)
	if !errors.Is(err, base) {
		t.Fatalf("wrapped error does not preserve base error: %v", err)
	}
	if !strings.Contains(err.Error(), "config.json") {
		t.Fatalf("wrapped error missing filename: %v", err)
	}
	if WrapReadError("config.json", nil) != nil {
		t.Fatal("WrapReadError(nil) should return nil")
	}
}
