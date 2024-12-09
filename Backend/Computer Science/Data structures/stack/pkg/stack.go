package stack

import "sync"

type (
	el[T any] struct {
		val  T
		next *el[T]
	}

	// LIFO implementation.
	Stack[T any] struct {
		mu    sync.Mutex
		first *el[T]
	}
)

// Creates new Stack with default capacity of 10.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.first = &el[T]{
		val:  elem,
		next: s.first,
	}
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if s.first == nil {
		return result, false
	}

	result = s.first.val
	s.first = s.first.next

	return result, true
}

func (s *Stack[T]) Len() int {
	s.mu.Lock()
	s.mu.Unlock()
	var (
		count int
		n     = s.first
	)
	for n != nil {
		count += 1
		n = n.next
	}

	return count
}
