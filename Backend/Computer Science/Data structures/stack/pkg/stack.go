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

// Push: Adding an element to the top of the stack.
// Asymptomatic : O(1)
func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.first = &element[T]{value: val, next: s.first}
	s.length++
}

// Pop: Removing an element from the top of the stack.
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

// Peek Looking at the top element of the stack without removing it.
// Asymptomatic : O(1)
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if s.first == nil {
		return result, false
	}

	return s.first.value, true
}

// Size: Getting the number of elements in the stack.
//Asymptomatic: O(1)
func (s *Stack[T]) Size() uint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

// Size: Getting the number of elements in the stack.
//Asymptomatic: O(1)
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Clear is method to truncate all elements in Stack.
// Asymptomatic: O(1)
func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.first = nil
	s.length = 0
}

// PopAll pops all elements and returns as slice.
// Asymptomatic: O(n)
func (s *Stack[T]) PopAll() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		result = make([]T, s.length)
		iter   = s.first
		idx    = 0
	)

	for iter != nil {
		result[idx] = iter.value
		iter = iter.next
		idx++
	}

	s.first = nil
	s.length = 0

	return result
}
