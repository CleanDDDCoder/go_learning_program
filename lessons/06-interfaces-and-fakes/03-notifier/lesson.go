package notifier

// Notifier sends a message to a user.
type Notifier interface {
	Notify(user string, message string) error
}

// SendWelcome sends a standard welcome notification.
func SendWelcome(notifier Notifier, user string) error {
	return notifier.Notify(user, "Welcome to Go Learning Program")
}
