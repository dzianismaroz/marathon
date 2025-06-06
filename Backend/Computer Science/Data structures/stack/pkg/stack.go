package stack

import (
	"fmt"
	"sync"
)

const intialCap = 10

// LIFO implementation.
type Stack[T any] struct {
	mu      sync.RWMutex
	content []T
}

// Creates new Stack with default capacity of 10.
func New[T any]() *Stack[T] {
	return &Stack[T]{content: make([]T, 0, intialCap)}
}

// Push: Adding an element to the top of the stack.
// Asymptotic : O(1).
func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.content = append(s.content, val)
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

	if len(s.content) == 0 {
		var result T

		return result, false
	}

	return s.content[len(s.content)-1], true
}

func (s *Stack[T]) Size() uint {
	return uint(len(s.content))
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.content) == 0
}

func (s *Stack[T]) PopAll() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]T, len(s.content))
	length := len(s.content) - 1

	for i, v := range s.content {
		result[length-i] = v
	}

	s.content = newSlice[T]()

	return result
}

func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.content = newSlice[T]()
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.content)
}

func newSlice[T any]() []T {
	return make([]T, intialCap)
}
