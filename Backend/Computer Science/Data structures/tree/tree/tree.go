package tree

type Tree[T any] struct {
	val   T
	left  *Tree[T]
	right *Tree[T]
}

func New[T any]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Traverse() {

}
