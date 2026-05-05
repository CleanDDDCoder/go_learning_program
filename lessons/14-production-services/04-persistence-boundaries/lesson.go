package main

import (
	"errors"
	"fmt"
	"sync"
)

// ErrNotFound is returned when a resource is not found
var ErrNotFound = errors.New("not found")

// Item represents a storable item
type Item struct {
	ID   string
	Name string
}

// Store defines the persistence interface
type Store interface {
	Create(item Item) error
	Get(id string) (Item, error)
	Update(item Item) error
	Delete(id string) error
}

// MemoryStore is an in-memory implementation
type MemoryStore struct {
	mu    sync.RWMutex
	items map[string]Item
}

// NewMemoryStore creates a new in-memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		items: make(map[string]Item),
	}
}

func (s *MemoryStore) Create(item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[item.ID]; ok {
		return errors.New("already exists")
	}
	s.items[item.ID] = item
	return nil
}

func (s *MemoryStore) Get(id string) (Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.items[id]
	if !ok {
		return Item{}, ErrNotFound
	}
	return item, nil
}

func (s *MemoryStore) Update(item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[item.ID]; !ok {
		return ErrNotFound
	}
	s.items[item.ID] = item
	return nil
}

func (s *MemoryStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return ErrNotFound
	}
	delete(s.items, id)
	return nil
}

func main() {
	store := NewMemoryStore()

	// Create
	store.Create(Item{ID: "1", Name: "Test"})

	// Get
	item, err := store.Get("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Got:", item)
	}

	// Update
	store.Update(Item{ID: "1", Name: "Updated"})
	item, _ = store.Get("1")
	fmt.Println("After update:", item)

	// Delete
	store.Delete("1")
	_, err = store.Get("1")
	fmt.Println("After delete:", err)
}