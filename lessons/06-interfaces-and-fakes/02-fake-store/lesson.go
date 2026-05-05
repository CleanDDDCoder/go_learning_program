package fake_store

// Store saves a key/value pair.
type Store interface {
	Save(key string, value string) error
}

// SaveGreeting writes a greeting message to the store.
func SaveGreeting(store Store, name string) error {
	return store.Save("greeting", "Hello, "+name+"!")
}
