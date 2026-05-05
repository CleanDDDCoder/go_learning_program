package sentinel_errors

import "errors"

// ErrNotFound is returned when an item is missing.
var ErrNotFound = errors.New("not found")

// FindName returns the name for id or ErrNotFound when id is missing.
func FindName(names map[int]string, id int) (string, error) {
	name, ok := names[id]
	if !ok {
		return "", ErrNotFound
	}
	return name, nil
}
