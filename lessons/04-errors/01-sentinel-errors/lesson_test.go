package sentinel_errors

import (
	"errors"
	"testing"
)

func TestFindName(t *testing.T) {
	names := map[int]string{1: "Ada"}

	got, err := FindName(names, 1)
	if err != nil {
		t.Fatalf("FindName existing err = %v, want nil", err)
	}
	if got != "Ada" {
		t.Fatalf("FindName existing = %q, want Ada", got)
	}

	_, err = FindName(names, 2)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("FindName missing err = %v, want ErrNotFound", err)
	}
}
