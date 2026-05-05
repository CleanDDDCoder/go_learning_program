//go:build ignore

package channels

// SumConcurrent returns the sum of numbers sent through the channel.
// It reads from the channel until it's closed, then returns the total.
func SumConcurrent(ch <-chan int) int {
	sum := 0
	for n := range ch {
		sum += n
	}
	return sum
}

// SendNumbers sends numbers 1 through n to the given channel and closes it.
func SendNumbers(ch chan<- int, n int) {
	defer close(ch)
	for i := 1; i <= n; i++ {
		ch <- i
	}
}
