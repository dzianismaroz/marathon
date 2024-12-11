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
		head   *element[T] // head of Stask to hold the latest element.
		length uint        // optimal way for Stack to track its length.
	}
)

var genericPools = sync.Map{} // Map to store pools by type.

// typedPool provides a type-safe wrapper for sync.Pool to manage generic Stack instances.
func typedPool[T any]() *sync.Pool {
	pool, _ := genericPools.LoadOrStore(new(T), &sync.Pool{
		New: func() any { return New[T]() },
	})
	return pool.(*sync.Pool)
}

// Creates new Stack instance.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
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

	s.head = &element[T]{value: val, next: s.head}
	s.length++
}

// Pop: Removing an element from the top of the stack.
// Asymptotic : O(1).
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
// Asymptotic : O(1).
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result T

	if s.head == nil {
		return result, false
	}

	return s.head.value, true
}

// Size: Getting the number of elements in the stack.
// Asymptotic: O(1).
func (s *Stack[T]) Size() uint {
	return s.length
}

// Size: resolves stack emptiness.
// Asymptotic: O(1).
func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
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
// Asymptotic: O(n).
func (s *Stack[T]) PopAll() []T {

	idx := 0

	s.mu.Lock() // --- critical section starts here ---.

	result := make([]T, s.length)
	head := s.head

	for head != nil {
		result[idx] = head.value
		head = head.next
		idx++
	}

	s.length = 0

	s.mu.Unlock() // -- critical section release ---.

	return result
}
