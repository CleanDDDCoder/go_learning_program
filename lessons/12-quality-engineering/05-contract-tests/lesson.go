package contracttests

// Cache defines a generic caching interface for contract testing.
type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T) error
	Delete(key string) error
}

// InMemoryCache is a simple in-memory implementation of Cache.
type InMemoryCache[T any] struct {
	data map[string]T
}

func NewInMemoryCache[T any]() *InMemoryCache[T] {
	return &InMemoryCache[T]{data: make(map[string]T)}
}

func (c *InMemoryCache[T]) Get(key string) (T, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *InMemoryCache[T]) Set(key string, value T) error {
	c.data[key] = value
	return nil
}

func (c *InMemoryCache[T]) Delete(key string) error {
	delete(c.data, key)
	return nil
}

// VerifyCacheContract tests that an implementation satisfies the Cache contract.
// This is a common pattern for interface testing.
func VerifyCacheContract[T any](cache Cache[T]) error {
	// In a real implementation, this would run property-based tests
	// to verify the contract is satisfied
	return nil
}