package fake_store

import (
	"errors"
	"testing"
)

type fakeStore struct {
	key   string
	value string
	err   error
}

func (store *fakeStore) Save(key string, value string) error {
	store.key = key
	store.value = value
	return store.err
}

func TestSaveGreetingStoresMessage(t *testing.T) {
	store := &fakeStore{}

	err := SaveGreeting(store, "Ada")
	if err != nil {
		t.Fatalf("SaveGreeting err = %v, want nil", err)
	}
	if store.key != "greeting" || store.value != "Hello, Ada!" {
		t.Fatalf("stored %q=%q, want greeting=Hello, Ada!", store.key, store.value)
	}
}

func TestSaveGreetingReturnsStoreError(t *testing.T) {
	want := errors.New("disk full")
	store := &fakeStore{err: want}

	err := SaveGreeting(store, "Ada")
	if !errors.Is(err, want) {
		t.Fatalf("SaveGreeting err = %v, want %v", err, want)
	}
}
