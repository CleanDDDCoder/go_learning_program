package retry

import (
	"errors"
	"testing"
)

func TestRetrySucceedsEventually(t *testing.T) {
	failures := 0
	err := Retry(3, func() error {
		failures++
		if failures < 2 {
			return errors.New("try again")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Retry returned err = %v, want nil", err)
	}
	if failures != 2 {
		t.Fatalf("attempts = %d, want 2", failures)
	}
}

func TestRetryReturnsLastError(t *testing.T) {
	want := errors.New("still failing")
	err := Retry(2, func() error {
		return want
	})
	if !errors.Is(err, want) {
		t.Fatalf("Retry err = %v, want %v", err, want)
	}
}
