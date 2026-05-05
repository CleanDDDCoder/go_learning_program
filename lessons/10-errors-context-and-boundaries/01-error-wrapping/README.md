# Error Wrapping and Inspection

This lesson teaches you how to:
- Wrap errors with context using `fmt.Errorf` and the `%w` verb
- Inspect wrapped errors using `errors.Is` and `errors.As`
- Preserve the underlying error chain for debugging

## Exercises

1. Implement a function that wraps errors with context
2. Use `errors.Is` and `errors.As` to inspect error chains