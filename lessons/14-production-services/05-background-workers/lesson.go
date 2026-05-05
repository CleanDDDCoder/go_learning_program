package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Job represents a unit of work
type Job struct {
	ID   int
	Data string
}

func main() {
	numWorkers := 3
	jobs := make(chan Job, 10)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d: %s\n", workerID, job.ID, job.Data)
			}
			fmt.Printf("Worker %d done\n", workerID)
		}(i)
	}

	// Send jobs
	for i := 1; i <= 5; i++ {
		jobs <- Job{ID: i, Data: fmt.Sprintf("task-%d", i)}
	}
	close(jobs)

	// Wait for signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutdown signal received")
	}()

	wg.Wait()
	fmt.Println("All jobs complete")
}