package queue

import "sync"

type (
	el[T any] struct {
		val  T
		next *el[T]
	}

	Queue[T any] struct {
		mu   sync.Mutex
		last *el[T]
	}
)

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(val T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		q.last = &el[T]{val: val}
		return
	}

	q.last = &el[T]{val: val, next: q.last}

	return
}

func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		var zero T
		return zero, false
	}

	val := q.last.val
	q.last = q.last.next

	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		var zero T
		return zero, false
	}

	return q.last.val, true
}

func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	var (
		count int
		p     = q.last
	)

	for p != nil {
		count += 1
		p = p.next
	}

	return count
}
