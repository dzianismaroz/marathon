package queue

import (
	"sync"
)

type Queue[T any] struct {
	mu      sync.RWMutex
	content []T
}

// Creates new Queue with default capacity of 10.
func New[T any]() *Queue[T] {
	return &Queue[T]{content: make([]T, 0, 10)}
}

// Asymptotic: O(1)
func (q *Queue[T]) Push(val T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.content = append(q.content, val)
}

// Asymptotic: O(1)
func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	switch {
	case len(q.content) == 0:
		var zero T
		return zero, false
	case len(q.content) == 1:
		result := q.content[0]
		q.content = make([]T, 0, 10)
		return result, true
	default:
		result := q.content[0]
		q.content = q.content[1:]
		return result, true
	}
}

// Asymptotic: O(1)
func (q *Queue[T]) Peek() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.content) == 0 {
		var zero T
		return zero, false
	}

	return q.content[0], true
}

// Asymptotic: O(1)
func (q *Queue[T]) Size() uint {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return uint(len(q.content))
}

func (q *Queue[T]) PopAll() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	result := make([]T, len(q.content))
	copy(result, q.content)
	q.content = make([]T, 0, 10)
	return result
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Clear is method to truncate all elements in Queue.
// Asymptotic: O(1)
func (q *Queue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.content = make([]T, 0, 10)
}
