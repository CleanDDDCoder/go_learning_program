//go:build ignore

package filtering

// EvenNumbers returns the even numbers from values in their original order.
func EvenNumbers(values []int) []int {
	var evens []int
	for _, value := range values {
		if value%2 == 0 {
			evens = append(evens, value)
		}
	}
	return evens
}
