package tree

type (
	node[T comparable] struct {
		val   T
		left  *node[T] // less than val
		right *node[T] // greater than val.
	}

	Tree[T comparable] struct {
		root     *node[T] // root tree node.
		nodesCol uint     // total count of nodes.
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
