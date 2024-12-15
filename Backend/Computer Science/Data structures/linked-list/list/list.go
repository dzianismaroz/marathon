package list

import "sync"

type (
	node[T any] struct {
		val  T
		next *node[T]
	}

	LinkedList[T any] struct {
		mu    sync.RWMutex
		first *node[T]
	}
)

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	next := &node[T]{val: v}

	n := l.first

	for n != nil {
		n = n.next
	}

	n = next
}
