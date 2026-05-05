package worker_pools

import (
	"sync"
)

func ProcessJobs(jobs <-chan int, numWorkers int) []int {
	results := []int{}
	var wg sync.WaitGroup

	// Mutex to protect results slice
	var mu sync.Mutex

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				mu.Lock()
				results = append(results, job)
				mu.Unlock()
			}
		}()
	}

	// Close jobs channel signals workers to exit
	// (caller should close jobs after sending all jobs)

	// Wait for all workers to complete
	wg.Wait()

	return results
}
