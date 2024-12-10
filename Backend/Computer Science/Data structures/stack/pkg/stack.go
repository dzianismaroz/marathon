package stack

import "sync"

type (

	// element is wrapper to hold stored value and link to next element.
	element[T any] struct {
		value T
		next  *element[T]
	}

	// LIFO implementation.
	Stack[T any] struct {
		mu     sync.RWMutex
		first  *element[T] // head of Stask.
		length uint        // optimal way for Stack to track its length.
	}
)

// Creates new Stack with default capacity of 10.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Asymptomatic : O(1)
func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.first = &element[T]{value: val, next: s.first}
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

	result, s.first = s.first.value, s.first.next
	s.length--

	return result, true
}

//Asymptomatic: O(1)
func (s *Stack[T]) Len() uint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}
