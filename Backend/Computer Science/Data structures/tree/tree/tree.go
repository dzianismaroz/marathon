package tree

import (
	"sync"
)

type (
	ordered interface {
		~int | ~float32 | ~float64 | ~uint | ~string
	}
	node[T ordered] struct {
		val   T
		left  *node[T] // less than val
		right *node[T] // greater than val.
	}

	Tree[T ordered] struct {
		mu       sync.RWMutex
		root     *node[T] // root tree node.
		nodesCol uint     // total count of nodes.
	}
)

func newNode[T ordered](elem T) *node[T] {
	return &node[T]{val: elem}
}

func New[T ordered]() *Tree[T] {
	return &Tree[T]{}
}

func (n *node[T]) add(elem T) bool {
	switch {
	case n.val < elem:
		if n.left == nil {
			n.left = newNode(elem)
			return true
		}

		return n.left.add(elem)
	case n.val > elem:
		if n.right == nil {
			n.right = newNode(elem)
			return true
		}
		return n.right.add(elem)
	}
	return false
}

func (t *Tree[T]) Add(elem T) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.root == nil {
		t.root = newNode(elem)
		return
	}

	if t.root.add(elem) {
		t.nodesCol++
	}
}

func (t *Tree[T]) Size() uint {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.nodesCol
}

func (n *node[T]) sorted() []T {
	var sorted []T
	if n.left != nil {
		sorted = append(sorted, n.left.sorted()...)
	}
	sorted = append(sorted, n.val)
	if n.right != nil {
		sorted = append(sorted, n.right.sorted()...)
	}
	return sorted
}

func (t *Tree[T]) SortedDesc() []T {
	var result []T
	if t.root == nil {
		return result
	}

	result = append(result, t.root.sorted()...)
	return result
}

func (t *Tree[T]) SortedAsc() []T {
	desc := t.SortedDesc()
	result := make([]T, len(desc))
	length := len(desc) - 1
	for i := 0; i <= length; i++ {
		result[i] = desc[length-i]
	}

	return result
}
