package main

import (
	"sync"
)

// MetricType represents the type of metric
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
)

// Metric represents a basic metric
type Metric struct {
	Name      string
	Type      MetricType
	Help      string
	mu        sync.RWMutex
	counter   uint64
	gauge     float64
	histogram []float64
}

// Counter implements an incrementing counter metric
type Counter struct {
	Metric
}

// NewCounter creates a new counter metric
func NewCounter(name, help string) *Counter {
	return &Counter{
		Metric: Metric{
			Name: name,
			Type: MetricTypeCounter,
			Help: help,
		},
	}
}

// Inc increments the counter by the given value
func (c *Counter) Inc(delta uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter += delta
}

// Value returns the current counter value
func (c *Counter) Value() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}

// Gauge implements a gauge metric (can go up and down)
type Gauge struct {
	Metric
}

// NewGauge creates a new gauge metric
func NewGauge(name, help string) *Gauge {
	return &Gauge{
		Metric: Metric{
			Name: name,
			Type: MetricTypeGauge,
			Help: help,
	},
	}
}

// Set sets the gauge to the given value
func (g *Gauge) Set(value float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.gauge = value
}

// Value returns the current gauge value
func (g *Gauge) Value() float64 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.gauge
}

// Inc increments the gauge by the given value
func (g *Gauge) Inc(delta float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.gauge += delta
}

// Dec decrements the gauge by the given value
func (g *Gauge) Dec(delta float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.gauge -= delta
}

// Histogram implements a histogram metric
type Histogram struct {
	Metric
	Buckets []float64
}

// NewHistogram creates a new histogram metric
func NewHistogram(name, help string, buckets []float64) *Histogram {
	if buckets == nil {
		buckets = []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}
	}
	return &Histogram{
		Metric: Metric{
			Name: name,
			Type: MetricTypeHistogram,
			Help: help,
		},
		Buckets: buckets,
	}
}

// Observe records an observation in the histogram
func (h *Histogram) Observe(value float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.histogram = append(h.histogram, value)
}

// Values returns all observed values
func (h *Histogram) Values() []float64 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	result := make([]float64, len(h.histogram))
	copy(result, h.histogram)
	return result
}

// Registry holds a collection of metrics
type Registry struct {
	mu      sync.RWMutex
	metrics map[string]interface{}
}

// NewRegistry creates a new metric registry
func NewRegistry() *Registry {
	return &Registry{
		metrics: make(map[string]interface{}),
	}
}

// Register registers a metric in the registry
func (r *Registry) Register(name string, metric interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.metrics[name] = metric
}

// GetCounter retrieves a counter from the registry
func (r *Registry) GetCounter(name string) (*Counter, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.metrics[name].(*Counter)
	return c, ok
}

// GetGauge retrieves a gauge from the registry
func (r *Registry) GetGauge(name string) (*Gauge, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	g, ok := r.metrics[name].(*Gauge)
	return g, ok
}

// GetHistogram retrieves a histogram from the registry
func (r *Registry) GetHistogram(name string) (*Histogram, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	h, ok := r.metrics[name].(*Histogram)
	return h, ok
}

func main() {}