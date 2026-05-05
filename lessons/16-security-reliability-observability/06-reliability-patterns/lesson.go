package main

import (
	"context"
	"errors"
	"time"
)

// ErrCircuitOpen is returned when the circuit breaker is open
var ErrCircuitOpen = errors.New("circuit breaker is open")

// ErrTimeout is returned when an operation times out
var ErrTimeout = errors.New("operation timed out")

// CircuitBreakerState represents the state of a circuit breaker
type CircuitBreakerState int

const (
	StateClosed CircuitBreakerState = iota
	StateOpen
	StateHalfOpen
)

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	failures         int
	successes        int
	state            CircuitBreakerState
	failureThreshold int
	successThreshold int
	timeout          time.Duration
	lastFailureTime  time.Time
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(failureThreshold, successThreshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		timeout:          timeout,
		state:            StateClosed,
	}
}

// Execute runs the given function if the circuit is closed
func (cb *CircuitBreaker) Execute(fn func() error) error {
	if cb.state == StateOpen {
		// Check if we should transition to half-open
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
			cb.successes = 0
		} else {
			return ErrCircuitOpen
		}
	}

	err := fn()

	if err != nil {
		cb.recordFailure()
	} else {
		cb.recordSuccess()
	}

	return err
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failures++
	cb.lastFailureTime = time.Now()

	if cb.state == StateHalfOpen {
		cb.state = StateOpen
	} else if cb.failures >= cb.failureThreshold {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.successes++

	if cb.state == StateHalfOpen {
		if cb.successes >= cb.successThreshold {
			cb.state = StateClosed
			cb.failures = 0
			cb.successes = 0
		}
	} else {
		cb.failures = 0
	}
}

// State returns the current state of the circuit breaker
func (cb *CircuitBreaker) State() CircuitBreakerState {
	return cb.state
}

// RetryConfig holds configuration for retry logic
type RetryConfig struct {
	MaxAttempts int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	BackoffMultiplier float64
}

// DefaultRetryConfig returns a default retry configuration
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts:      3,
		InitialDelay:     100 * time.Millisecond,
		MaxDelay:         5 * time.Second,
		BackoffMultiplier: 2.0,
	}
}

// Retry executes a function with exponential backoff retry
func Retry(ctx context.Context, config RetryConfig, fn func() error) error {
	var lastErr error
	delay := config.InitialDelay

	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}

		lastErr = err

		// Don't retry on the last attempt
		if attempt < config.MaxAttempts-1 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}

			// Calculate next delay with exponential backoff
			delay = time.Duration(float64(delay) * config.BackoffMultiplier)
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
		}
	}

	return lastErr
}

// TimeoutWrapper wraps a function with a timeout
func TimeoutWrapper(timeout time.Duration, fn func() error) func() error {
	return func() error {
		done := make(chan error, 1)
		
		go func() {
			done <- fn()
		}()

		select {
		case err := <-done:
			return err
		case <-time.After(timeout):
			return ErrTimeout
		}
	}
}

// RateLimiter implements a simple rate limiter
type RateLimiter struct {
	tokens    float64
	maxTokens float64
	refillRate float64
	lastRefill time.Time
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(maxTokens float64, refillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:    maxTokens,
		maxTokens: maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if an operation is allowed and consumes a token if available
func (rl *RateLimiter) Allow() bool {
	rl.refill()

	if rl.tokens >= 1 {
		rl.tokens -= 1
		return true
	}
	return false
}

func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()
	rl.tokens += elapsed * rl.refillRate
	if rl.tokens > rl.maxTokens {
		rl.tokens = rl.maxTokens
	}
	rl.lastRefill = now
}

func main() {}