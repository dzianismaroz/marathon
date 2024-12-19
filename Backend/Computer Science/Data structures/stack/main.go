package main

import (
	"fmt"

	stack "dzianismaroz.github.com/marathon/stack/pkg"
)

func main() {
	stk := stack.New[int]()
	stk.Push(10)
	stk.Push(9)
	stk.Push(8)
	fmt.Println("stack after adding 3 elements Size:", stk.Size())
	fmt.Println("stack after adding 10, 9, 8:", stk)

	peeked, _ := stk.Peek()
	fmt.Println("peeking stack:", peeked)
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println("stack size after popping 3 elements:", stk.Size())

	fmt.Println(stk.Pop())
	fmt.Println("stack after popping empty stack Size:", stk.Size())
}
