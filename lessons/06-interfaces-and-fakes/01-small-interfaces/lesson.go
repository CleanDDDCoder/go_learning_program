package small_interfaces

// Greeter returns a greeting for a name.
type Greeter interface {
	Greet(name string) string
}

// Welcome uses a Greeter dependency to build a welcome message.
func Welcome(greeter Greeter, name string) string {
	return greeter.Greet(name)
}
