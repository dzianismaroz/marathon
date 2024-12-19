package list

import (
	"errors"
	"sync"
)

type (
	node[T comparable] struct {
		val  T
		next *node[T]
	}

	LinkedList[T comparable] struct {
		mu   sync.RWMutex
		head *node[T]
		size uint
	}
)

// New creates a new list from the given items.
func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func BuildFrom[T comparable](items []T) *LinkedList[T] {
	l := New[T]()

	for _, item := range items {
		l.Append(item)
	}

	return l
}

// Append adds item to the end of the list.
// Asymptotic: O(n)
func (l *LinkedList[T]) Append(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.size == 0 {
		l.head = &node[T]{val: v}
		l.size++

		return
	}

	last := l.head

	for last.next != nil {
		last = last.next
	}

	last.next = &node[T]{val: v}
	l.size++
}

// First returns the first item of the list if presented.
// Asymptotic: O(1)
func (l *LinkedList[T]) First() (T, bool) {
	if l.size == 0 {
		var zero T

		return zero, false
	}

	return l.head.val, true
}

// Size returns the size of the list.
// Asymptotic: O(1)
func (l *LinkedList[T]) Size() uint {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.size
}

// IndexOf returns the index of item in the list.
// Asymptotic: O(n)
func (l *LinkedList[T]) IndexOf(item T) (uint, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var i uint

	for ; i < l.size; i++ {
		if l.head.val == item {
			return i, true
		}

		l.head = l.head.next
	}

	return 0, false
}

// InsertAt inserts item at index idx of the list.
// Asymptotic: O(n)
func (l *LinkedList[T]) InsertAt(idx uint, item T) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if idx > l.size {
		return errors.New("index out of range")
	}

	if idx == 0 {
		prev := l.head
		l.head = &node[T]{val: item, next: prev}
		l.size++
		return nil
	}

	n := l.head
	var count uint
	previous := n

	for count < idx {
		previous = n
		count++
		n = n.next
	}

	previous.next = &node[T]{val: item, next: n}
	l.size++

	return nil
}

// RemoveAt removes item at index idx of the list.
// Asymptotic: O(n)
func (l *LinkedList[T]) RemoveAt(idx uint) (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if idx >= l.size {
		var zero T
		return zero, errors.New("index out of range")
	}

	if idx == 0 {
		result := l.head.val
		l.head = l.head.next
		l.size--
		return result, nil
	}

	node := l.head
	var count uint
	previous := node

	for count < idx {
		previous = node
		count++
		node = node.next
	}

	toRemove := node
	previous.next = toRemove.next
	l.size--

	return toRemove.val, nil
}

// Items returns the items slice of the list.
// Asymptotic: O(n)
func (l *LinkedList[T]) Items() []T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	result := make([]T, 0, l.size)
	next := l.head

	for next != nil {
		result = append(result, next.val)
		next = next.next
	}

	return result
}
