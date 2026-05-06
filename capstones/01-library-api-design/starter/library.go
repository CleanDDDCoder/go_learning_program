package library

import "errors"

// ErrNotFound is returned when a requested record does not exist.
var ErrNotFound = errors.New("record not found")

// Store provides the public API learners refine during the capstone.
type Store struct {
	records map[string]string
}

// NewStore creates an empty Store.
func NewStore() *Store {
	return &Store{records: map[string]string{}}
}

// Put stores a value by key.
func (s *Store) Put(key string, value string) error {
	// TODO: Validate the public API contract and document compatibility choices.
	s.records[key] = value
	return nil
}

// Get returns the stored value for key.
func (s *Store) Get(key string) (string, error) {
	value, ok := s.records[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
