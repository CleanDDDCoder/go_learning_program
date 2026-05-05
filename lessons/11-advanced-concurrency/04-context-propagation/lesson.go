package context_propagation

import "context"

// TODO: Implement a function that spawns multiple goroutines
// that all respect a shared context.
// - Spawn `numWorkers` goroutines.
// - Each worker should check for context cancellation periodically.
// - Return true if all workers complete successfully.
// - Return false if the context is cancelled.

func RunWithContext(ctx context.Context, numWorkers int) bool {
	// Your implementation here
	return true
}