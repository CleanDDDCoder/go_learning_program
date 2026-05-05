//go:build ignore

package context_cancellation

import "context"

// CollectUntilDone gathers values until input closes or ctx is canceled.
func CollectUntilDone(ctx context.Context, input <-chan int) []int {
	var values []int
	for {
		select {
		case <-ctx.Done():
			return values
		case value, ok := <-input:
			if !ok {
				return values
			}
			values = append(values, value)
		}
	}
}
