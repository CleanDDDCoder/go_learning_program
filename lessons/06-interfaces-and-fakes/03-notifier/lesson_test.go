package notifier

import "testing"

type fakeNotifier struct {
	user    string
	message string
}

func (fake *fakeNotifier) Notify(user string, message string) error {
	fake.user = user
	fake.message = message
	return nil
}

func TestSendWelcomeNotifiesUser(t *testing.T) {
	fake := &fakeNotifier{}

	err := SendWelcome(fake, "sam")
	if err != nil {
		t.Fatalf("SendWelcome err = %v, want nil", err)
	}
	if fake.user != "sam" {
		t.Fatalf("fake user = %q, want sam", fake.user)
	}
	if fake.message != "Welcome to Go Learning Program" {
		t.Fatalf("fake message = %q, want welcome message", fake.message)
	}
}
