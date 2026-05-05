//go:build ignore

package constructors

// Counter stores a mutable count.
type Counter struct {
	Value int
}

// NewCounter creates a counter with the provided starting value.
func NewCounter(start int) Counter {
	return Counter{Value: start}
}

// Increment increases the counter by one.
func (counter *Counter) Increment() {
	counter.Value++
}
