package stack

import "sync"

type (
	el[T any] struct {
		val  T
		next *el[T]
	}

	// LIFO implementation.
	Stack[T any] struct {
		mu     sync.RWMutex
		first  *el[T]
		length uint // optimal way for Stack to track its length.
	}
)

// Creates new Stack with default capacity of 10.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Asymptomatic : O(1)
func (s *Stack[T]) Push(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.first = &el[T]{ // O(1)
		val:  elem,
		next: s.first,
	}

	s.length++
}

// Asymptomatic : O(1)
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if s.first == nil {
		return result, false
	}

	result = s.first.val
	s.first = s.first.next

	s.length--

	return result, true
}

//Asymptomatic: O(1)
func (s *Stack[T]) Len() uint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}
