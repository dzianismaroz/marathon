package tree

type (
	node[T comparable] struct {
		val   T
		left  *node[T]
		right *node[T]
	}

	Tree[T comparable] struct {
		root *node[T]
	}
)

func New[T comparable]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Add(elem T) {
	if t.root == nil {
		t.root = &node[T]{val: elem}

		return
	}

}

func (t *Tree[T]) addRecursive(n *node[T], elem T) {
	if n.val < elem {

	}

}
