package stack

type LIFO interface {
	Push(any)
	Pop() (any, bool)
	Peek() (any, bool)
	Size() uint
	Clear()
	PopAll() []any
}
