package methods

import "testing"

func TestRectangleMethods(t *testing.T) {
	rect := Rectangle{Width: 4, Height: 7}

	if got := rect.Area(); got != 28 {
		t.Fatalf("Area() = %d, want 28", got)
	}
	if got := rect.Perimeter(); got != 22 {
		t.Fatalf("Perimeter() = %d, want 22", got)
	}
}
