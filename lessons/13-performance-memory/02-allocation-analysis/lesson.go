package allocation

import "strings"

// FilterEmpty filters out empty strings from the input and returns the remaining strings.
//
// TODO: This implementation allocates a new slice on every call.
// TODO: Refactor to pre-allocate the result slice with capacity to avoid repeated allocations.
func FilterEmpty(items []string) []string {
	result := []string{}
	for _, s := range items {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// Concat joins strings with a separator.
//
// TODO: This implementation uses strings.Join which allocates.
// TODO: Refactor to use strings.Builder with pre-allocated capacity.
func Concat(items []string, sep string) string {
	return strings.Join(items, sep)
}