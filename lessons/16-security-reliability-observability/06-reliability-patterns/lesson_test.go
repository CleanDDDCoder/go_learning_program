package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	t.Run("starts in closed state", func(t *testing.T) {
		cb := NewCircuitBreaker(3, 2, 10*time.Second)
		if cb.State() != StateClosed {
			t.Errorf("expected StateClosed, got %v", cb.State())
		}
	})

	t.Run("opens after threshold failures", func(t *testing.T) {
		cb := NewCircuitBreaker(3, 2, 10*time.Second)
		
		for i := 0; i < 3; i++ {
			cb.Execute(func() error { return errors.New("fail") })
		}
		
		if cb.State() != StateOpen {
			t.Errorf("expected StateOpen, got %v", cb.State())
		}
	})

	t.Run("rejects requests when open", func(t *testing.T) {
		cb := NewCircuitBreaker(1, 2, 10*time.Second)
		cb.Execute(func() error { return errors.New("fail") })
		
		err := cb.Execute(func() error { return nil })
		if err != ErrCircuitOpen {
			t.Errorf("expected ErrCircuitOpen, got %v", err)
		}
	})

	t.Run("half-open after timeout", func(t *testing.T) {
		cb := NewCircuitBreaker(1, 2, 50*time.Millisecond)
		cb.Execute(func() error { return errors.New("fail") })
		
		time.Sleep(60 * time.Millisecond)
		cb.Execute(func() error { return nil }) // Should not return ErrCircuitOpen
		
		if cb.State() != StateHalfOpen {
			t.Errorf("expected StateHalfOpen, got %v", cb.State())
		}
	})
}

func TestRetry(t *testing.T) {
	t.Run("succeeds on first attempt", func(t *testing.T) {
		config := DefaultRetryConfig()
		ctx := context.Background()
		
		err := Retry(ctx, config, func() error { return nil })
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("retries on failure", func(t *testing.T) {
		config := DefaultRetryConfig()
		ctx := context.Background()
		attempts := 0
		
		err := Retry(ctx, config, func() error {
			attempts++
			if attempts < 3 {
				return errors.New("fail")
			}
			return nil
		})
		
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if attempts != 3 {
			t.Errorf("expected 3 attempts, got %d", attempts)
		}
	})

	t.Run("returns last error after max attempts", func(t *testing.T) {
		config := DefaultRetryConfig()
		ctx := context.Background()
		
		err := Retry(ctx, config, func() error { return errors.New("persistent failure") })
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestTimeoutWrapper(t *testing.T) {
	t.Run("completes within timeout", func(t *testing.T) {
		fn := TimeoutWrapper(100*time.Millisecond, func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		})
		
		err := fn()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("times out", func(t *testing.T) {
		fn := TimeoutWrapper(50*time.Millisecond, func() error {
			time.Sleep(100 * time.Millisecond)
			return nil
		})
		
		err := fn()
		if err != ErrTimeout {
			t.Errorf("expected ErrTimeout, got %v", err)
		}
	})
}

func TestRateLimiter(t *testing.T) {
	t.Run("allowsrequests up to limit", func(t *testing.T) {
		rl := NewRateLimiter(5, 1) // 5 tokens, refill at 1 per second
		
		for i := 0; i < 5; i++ {
			if !rl.Allow() {
				t.Errorf("expected Allow()=true on attempt %d", i+1)
			}
		}
		
		if rl.Allow() {
			t.Error("expected Allow()=false when exhausted")
		}
	})

	t.Run("refills over time", func(t *testing.T) {
		rl := NewRateLimiter(1, 100) // 1 token, refill at 100 per second
		
		rl.Allow() // Use the only token
		
		time.Sleep(20 * time.Millisecond) // Should refill ~2 tokens
		
		if !rl.Allow() {
			t.Error("expected Allow()=true after refill")
		}
	})
}