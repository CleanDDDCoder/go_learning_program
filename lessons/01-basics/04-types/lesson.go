package types

import "strconv"

// IntToString converts an integer to its string representation.
// Examples: 42 -> "42", 0 -> "0", -123 -> "-123"
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// StringToInt converts a string to an integer.
// Return 0 if the string is not a valid integer.
func StringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}