package queue

import "sync"

type (
	el[T any] struct {
		val  T
		next *el[T]
	}

	Queue[T any] struct {
		mu     sync.RWMutex
		last   *el[T]
		length uint // Queue is able to track it's own length
	}
)

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Asymptomatic: O(1)
func (q *Queue[T]) Push(val T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		q.last = &el[T]{val: val}
		return
	}

	q.last = &el[T]{val: val, next: q.last}
	q.length++

	return
}

// Asymptomatic: O(1)
func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		var zero T
		return zero, false
	}

	val := q.last.val
	q.last = q.last.next
	q.length--

	return val, true
}

// Asymptomatic: O(1)
func (q *Queue[T]) Peek() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		var zero T
		return zero, false
	}

	return q.last.val, true
}

// Asymptomatic: O(1)
func (q *Queue[T]) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return int(q.length)
}
