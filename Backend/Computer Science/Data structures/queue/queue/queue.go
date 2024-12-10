package queue

import "sync"

type (
	el[T any] struct {
		val  T
		next *el[T]
	}

	Queue[T any] struct {
		mu     sync.RWMutex
		first  *el[T]
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

	newEl := &el[T]{val: val}

	if q.last != nil {
		q.last.next = newEl
	} else {
		// If the queue is empty, first and last point to the new element
		q.first = newEl
	}

	q.last = newEl
	q.length++
}

// Asymptomatic: O(1)
func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.first == nil {
		var zero T
		return zero, false
	}

	val := q.first.val
	q.first = q.first.next
	if q.first == nil {
		// If the queue is now empty, reset the last pointer
		q.last = nil
	}
	q.length--

	return val, true
}

// Asymptomatic: O(1)
func (q *Queue[T]) Peek() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.first == nil {
		var zero T
		return zero, false
	}

	return q.first.val, true
}

// Asymptomatic: O(1)
func (q *Queue[T]) Len() uint {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.length
}
