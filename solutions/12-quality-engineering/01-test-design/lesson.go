//go:build ignore

// Package testdesign provides lessons on test design patterns.
package testdesign

// Sum returns the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

// Max returns the larger of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Reverse reverses a slice of strings.
func Reverse(s []string) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}