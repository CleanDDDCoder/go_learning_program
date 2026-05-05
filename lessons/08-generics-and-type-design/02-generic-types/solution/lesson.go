//go:build ignore

package main

// Stack is a generic stack that can hold any type
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// It returns the zero value of T if the stack is empty
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	println("Pop:", stack.Pop()) // 3
	println("Pop:", stack.Pop()) // 2
	println("Pop:", stack.Pop()) // 1
	println("Empty:", stack.IsEmpty()) // true
	
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	
	println("Pop:", strStack.Pop()) // world
	println("Pop:", strStack.Pop()) // hello
}