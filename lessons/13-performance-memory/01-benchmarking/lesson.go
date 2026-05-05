package benchmarking

import "testing"

// Sum returns the sum of all numbers in the slice.
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// BenchmarkSum benchmarks the Sum function.
func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}