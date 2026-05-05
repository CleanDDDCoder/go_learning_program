//go:build ignore

package retry

// Retry calls fn until it succeeds or attempts are exhausted.
func Retry(attempts int, fn func() error) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		if err := fn(); err != nil {
			lastErr = err
			continue
		}
		return nil
	}
	return lastErr
}
