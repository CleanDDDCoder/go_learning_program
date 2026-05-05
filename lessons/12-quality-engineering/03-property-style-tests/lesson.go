package propertytests

// Property-based testing demonstrates testing with randomized inputs
// to verify properties that should hold true across all valid inputs.

func ReverseReverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

func SumInts(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}