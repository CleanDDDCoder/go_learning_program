package race_debugging

// TODO: Fix the data race in this code.
// The current implementation has a race condition when
// multiple goroutines call Increment and Read simultaneously.
// Use sync.Mutex or sync.RWMutex to fix the race.

type Counter struct {
	// Add your synchronization primitive here
	value int
}

func (c *Counter) Increment() {
	// Your implementation here
	c.value++
}

func (c *Counter) Read() int {
	// Your implementation here
	return c.value
}
