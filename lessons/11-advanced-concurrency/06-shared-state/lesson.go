package shared_state

import (
	"sync"
	"sync/atomic"
)

// Part 1: Use sync.Mutex to protect shared state
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Part 2: Use atomic operations for simple counters
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Part 3: Use sync.RWMutex for read-heavy workloads
type RWCounter struct {
	mu    sync.RWMutex
	value int
}

func (c *RWCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *RWCounter) Read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
