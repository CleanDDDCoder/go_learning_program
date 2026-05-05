package main

import (
	"testing"
)

func TestStackInt(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after pushing items")
	}

	popped := stack.Pop()
	if popped != 3 {
		t.Errorf("Pop() = %d; want 3", popped)
	}

	popped = stack.Pop()
	if popped != 2 {
		t.Errorf("Pop() = %d; want 2", popped)
	}

	popped = stack.Pop()
	if popped != 1 {
		t.Errorf("Pop() = %d; want 1", popped)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all items")
	}
}

func TestStackString(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("hello")
	stack.Push("world")

	popped := stack.Pop()
	if popped != "world" {
		t.Errorf("Pop() = %q; want \"world\"", popped)
	}

	popped = stack.Pop()
	if popped != "hello" {
		t.Errorf("Pop() = %q; want \"hello\"", popped)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty")
	}
}

func TestStackEmptyPop(t *testing.T) {
	stack := NewStack[int]()

	// Popping from empty stack should return zero value
	popped := stack.Pop()
	if popped != 0 {
		t.Errorf("Pop() from empty stack = %d; want 0", popped)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should still be empty")
	}
}

func TestStackFloat(t *testing.T) {
	stack := NewStack[float64]()

	stack.Push(1.5)
	stack.Push(2.5)
	stack.Push(3.5)

	popped := stack.Pop()
	if popped != 3.5 {
		t.Errorf("Pop() = %v; want 3.5", popped)
	}

	popped = stack.Pop()
	if popped != 2.5 {
		t.Errorf("Pop() = %v; want 2.5", popped)
	}

	popped = stack.Pop()
	if popped != 1.5 {
		t.Errorf("Pop() = %v; want 1.5", popped)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty")
	}
}
