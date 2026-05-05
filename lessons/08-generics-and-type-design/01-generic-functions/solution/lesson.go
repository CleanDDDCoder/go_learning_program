//go:build ignore

package main

import "fmt"

// FizzBuzz returns "Fizz", "Buzz", "FizzBuzz", or the number as a string
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

// Number is a constraint that matches integers
type Number interface {
	int | int64 | float64
}

// Sum returns the sum of a slice of numbers
func Sum[T Number](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}

// Double returns twice the input value
func Double[T Number](v T) T {
	return v + v
}

func main() {
	fmt.Println("Generic Functions Lesson")
	fmt.Println("FizzBuzz(15):", FizzBuzz(15))
	fmt.Println("Sum of ints:", Sum([]int{1, 2, 3}))
	fmt.Println("Sum of floats:", Sum([]float64{1.5, 2.5}))
	fmt.Println("Double(7):", Double(7))
}