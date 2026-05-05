package constructors

import "testing"

func TestCounter(t *testing.T) {
	counter := NewCounter(10)
	if counter.Value != 10 {
		t.Fatalf("NewCounter(10).Value = %d, want 10", counter.Value)
	}

	counter.Increment()
	if counter.Value != 11 {
		t.Fatalf("Value after Increment() = %d, want 11", counter.Value)
	}
}
