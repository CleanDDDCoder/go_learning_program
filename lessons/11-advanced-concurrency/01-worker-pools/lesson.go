package worker_pools

// TODO: Implement a worker pool that processes jobs from a channel.
// - Start exactly `numWorkers` goroutines.
// - Each worker should read from the jobs channel until it's closed.
// - Process each job by adding its value to the results slice.
// - Use sync.WaitGroup to ensure all workers finish before returning.
// - Return the results in order.

func ProcessJobs(jobs <-chan int, numWorkers int) []int {
	// Your implementation here
	results := []int{}
	// TODO: Add your worker pool implementation
	return results
}