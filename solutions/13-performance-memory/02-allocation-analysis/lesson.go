package allocation

import "strings"

// FilterEmpty filters out empty strings from the input and returns the remaining strings.
// Optimized version: pre-allocates result slice with capacity.
func FilterEmpty(items []string) []string {
	// Pre-allocate with capacity to avoid repeated allocations
	result := make([]string, 0, len(items))
	for _, s := range items {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// Concat joins strings with a separator.
// Optimized version: uses strings.Builder with pre-allocated capacity.
func Concat(items []string, sep string) string {
	if len(items) == 0 {
		return ""
	}

	// Calculate total size needed
	n := len(sep) * (len(items) - 1)
	for _, s := range items {
		n += len(s)
	}

	// Use Builder with pre-allocated capacity
	var b strings.Builder
	b.Grow(n)
	for i, s := range items {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(s)
	}
	return b.String()
}