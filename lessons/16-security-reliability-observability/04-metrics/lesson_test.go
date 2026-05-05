package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("new counter starts at zero", func(t *testing.T) {
		c := NewCounter("test_counter", "A test counter")
		if c.Value() != 0 {
			t.Errorf("expected 0, got %d", c.Value())
		}
	})

	t.Run("inc increments counter", func(t *testing.T) {
		c := NewCounter("test_counter", "A test counter")
		c.Inc(5)
		if c.Value() != 5 {
			t.Errorf("expected 5, got %d", c.Value())
		}
		c.Inc(3)
		if c.Value() != 8 {
			t.Errorf("expected 8, got %d", c.Value())
		}
	})
}

func TestGauge(t *testing.T) {
	t.Run("new gauge starts at zero", func(t *testing.T) {
		g := NewGauge("test_gauge", "A test gauge")
		if g.Value() != 0 {
			t.Errorf("expected 0, got %f", g.Value())
		}
	})

	t.Run("set sets gauge value", func(t *testing.T) {
		g := NewGauge("test_gauge", "A test gauge")
		g.Set(42.5)
		if g.Value() != 42.5 {
			t.Errorf("expected 42.5, got %f", g.Value())
		}
	})

	t.Run("inc increments gauge", func(t *testing.T) {
		g := NewGauge("test_gauge", "A test gauge")
		g.Inc(10)
		if g.Value() != 10 {
			t.Errorf("expected 10, got %f", g.Value())
		}
	})

	t.Run("dec decrements gauge", func(t *testing.T) {
		g := NewGauge("test_gauge", "A test gauge")
		g.Set(10)
		g.Dec(3)
		if g.Value() != 7 {
			t.Errorf("expected 7, got %f", g.Value())
		}
	})
}

func TestHistogram(t *testing.T) {
	t.Run("observe records values", func(t *testing.T) {
		h := NewHistogram("test_histogram", "A test histogram", nil)
		h.Observe(0.5)
		h.Observe(1.0)
		h.Observe(1.5)

		values := h.Values()
		if len(values) != 3 {
			t.Errorf("expected 3 values, got %d", len(values))
		}
	})
}

func TestRegistry(t *testing.T) {
	t.Run("register and retrieve counter", func(t *testing.T) {
		reg := NewRegistry()
		c := NewCounter("test_counter", "A test counter")
		reg.Register("test_counter", c)

		retrieved, ok := reg.GetCounter("test_counter")
		if !ok {
			t.Error("expected to find counter")
		}
		if retrieved.Value() != 0 {
			t.Errorf("expected 0, got %d", retrieved.Value())
		}
	})

	t.Run("register and retrieve gauge", func(t *testing.T) {
		reg := NewRegistry()
		g := NewGauge("test_gauge", "A test gauge")
		reg.Register("test_gauge", g)

		retrieved, ok := reg.GetGauge("test_gauge")
		if !ok {
			t.Error("expected to find gauge")
		}
		retrieved.Set(100)
		if retrieved.Value() != 100 {
			t.Errorf("expected 100, got %f", retrieved.Value())
		}
	})

	t.Run("register and retrieve histogram", func(t *testing.T) {
		reg := NewRegistry()
		h := NewHistogram("test_histogram", "A test histogram", nil)
		reg.Register("test_histogram", h)

		retrieved, ok := reg.GetHistogram("test_histogram")
		if !ok {
			t.Error("expected to find histogram")
		}
		retrieved.Observe(0.1)
		if len(retrieved.Values()) != 1 {
			t.Errorf("expected 1 value, got %d", len(retrieved.Values()))
		}
	})
}

func TestConcurrentMetricAccess(t *testing.T) {
	t.Run("counter is thread-safe", func(t *testing.T) {
		c := NewCounter("concurrent_counter", "A concurrent counter")
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc(1)
			}()
		}
		wg.Wait()
		if c.Value() != 100 {
			t.Errorf("expected 100, got %d", c.Value())
		}
	})

	t.Run("gauge is thread-safe", func(t *testing.T) {
		g := NewGauge("concurrent_gauge", "A concurrent gauge")
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				g.Inc(1)
			}()
		}
		wg.Wait()
		if g.Value() != 100 {
			t.Errorf("expected 100, got %f", g.Value())
		}
	})
}