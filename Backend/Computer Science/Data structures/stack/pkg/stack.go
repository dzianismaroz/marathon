package stack

import (
	"fmt"
	"sync"
)

// LIFO implementation.
type Stack[T any] struct {
	mu      sync.RWMutex
	content []T
}

// Creates new Stack with default capacity of 10.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{content: make([]T, 0, 10)}
}

// Single provides a singleton Stack instance.
func Single[T any]() *Stack[T] {
	return typedPool[T]().Get().(*Stack[T])
}

// Release resets and returns the stack instance back to the pool.
func (s *Stack[T]) Release() {
	s.Clear() // Reset the stack before putting it back.
	typedPool[T]().Put(s)
}

// Push: Adding an element to the top of the stack.
// Asymptotic : O(1).
func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.content = append(s.content, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if len(s.content) == 0 {
		return result, false
	}

	result = s.content[len(s.content)-1]
	s.content = s.content[:len(s.content)-1]

	return result, true
}

func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result T
	if len(s.content) == 0 {
		return result, false
	}
	return s.content[len(s.content)-1], true
}

func (s *Stack[T]) Size() int {
	return len(s.content)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.content) == 0
}

func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.content = make([]T, 0, 10)
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.content)
}
