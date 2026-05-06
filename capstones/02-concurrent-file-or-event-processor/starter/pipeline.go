package capstone

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

// Event represents a unified event from any source.
type Event struct {
	ID       string
	Type     EventType
	Source   string
	Payload  []byte
	Received time.Time
}

// EventType categorizes events.
type EventType int

const (
	EventTypeFile EventType = iota
	EventTypeWebhook
	EventTypeMessage
)

// Processor processes events.
type Processor interface {
	Process(ctx context.Context, event Event) error
}

// Metrics holds processing statistics.
type Metrics struct {
	counters [3]int64
}

// RecordEvent records an event processed.
func (m *Metrics) RecordEvent(eventType EventType) {
	atomic.AddInt64(&m.counters[eventType], 1)
}

// ErrBackpressure indicates the pipeline is overwhelmed.
var ErrBackpressure = errors.New("backpressure: pipeline full")
