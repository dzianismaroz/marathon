package stack

import "sync"

// LIFO implementation.
type Stack[T comparable] struct {
	mu         sync.Mutex
	containter []T
}

// Creates new Stack with default capacity of 10.
func NewStack[T comparable]() *Stack[T] {
	container := make([]T, 0, 10)

	return &Stack[T]{containter: container}
}

func (s *Stack[T]) Push(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.containter = append(s.containter, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	last := len(s.containter) - 1
	if last < 0 {
		return result, false
	}

	result = s.containter[last]
	s.containter = s.containter[:last]

	return result, true
}

func (s *Stack[T]) Len() int {
	s.mu.Lock()
	s.mu.Unlock()

	return len(s.containter)
}
