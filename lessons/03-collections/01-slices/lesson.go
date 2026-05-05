package slices

// Sum returns the total of all numbers in values.
func Sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// Last returns the final item in values and reports whether one existed.
func Last(values []string) (string, bool) {
	if len(values) == 0 {
		return "", false
	}
	return values[len(values)-1], true
}
