package select_and_backpressure

import "context"

// TODO: Implement a function that merges two channels with backpressure.
// - Read from both input channels using select.
// - When both channels are done, close the output channel.
// - Implement bounded buffering to apply backpressure.
// - Return the merged channel.

func MergeWithBackpressure(ctx context.Context, ch1, ch2 <-chan int, bufferSize int) <-chan int {
	// Your implementation here
	return nil
}
