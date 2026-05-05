package success_failure

import (
	"errors"
	"testing"
)

func TestDivideSuccess(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide returned err = %v, want nil", err)
	}
	if got != 5 {
		t.Fatalf("Divide(10, 2) = %d, want 5", got)
	}
}

func TestDivideFailure(t *testing.T) {
	_, err := Divide(10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Divide(10, 0) err = %v, want ErrDivideByZero", err)
	}
}
