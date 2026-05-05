package small_interfaces

import "testing"

type recordingGreeter struct {
	calledWith string
}

func (greeter *recordingGreeter) Greet(name string) string {
	greeter.calledWith = name
	return "Welcome, " + name
}

func TestWelcomeUsesGreeter(t *testing.T) {
	fake := &recordingGreeter{}

	got := Welcome(fake, "Ada")
	if got != "Welcome, Ada" {
		t.Fatalf("Welcome() = %q, want Welcome, Ada", got)
	}
	if fake.calledWith != "Ada" {
		t.Fatalf("fake calledWith = %q, want Ada", fake.calledWith)
	}
}
