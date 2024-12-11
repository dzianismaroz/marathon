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
		head   *element[T] // head of Stask.
		length uint        // optimal way for Stack to track its length.
	}
)

// Creates new Stack instance.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push: Adding an element to the top of the stack.
// Asymptotic : O(1)
func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.head = &element[T]{value: val, next: s.head}
	s.length++
}

// Pop: Removing an element from the top of the stack.
// Asymptotic : O(1)
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if s.head == nil {
		return result, false
	}

	result, s.head = s.head.value, s.head.next
	s.length--

	return result, true
}

// Peek Looking at the top element of the stack without removing it.
// Asymptotic : O(1)
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result T

	if s.head == nil {
		return result, false
	}

	return s.head.value, true
}

// Size: Getting the number of elements in the stack.
//Asymptotic: O(1)
func (s *Stack[T]) Size() uint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

// Size: Getting the number of elements in the stack.
//Asymptotic: O(1)
func (s *Stack[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Size() == 0
}

// Clear is method to truncate all elements in Stack.
// Asymptotic: O(1)
func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.head = nil
	s.length = 0
}

// PopAll pops all elements and returns as slice.
// Asymptotic: O(n)
func (s *Stack[T]) PopAll() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]T, s.length)

	for s.head != nil {
		result = append(result, s.head.value)
		s.head = s.head.next
	}

	s.length = 0

	return result
}
