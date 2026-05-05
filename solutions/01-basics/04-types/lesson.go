//go:build ignore

package types

import "strconv"

// IntToString converts an integer to its string representation.
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// StringToInt converts a string to an integer.
func StringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}