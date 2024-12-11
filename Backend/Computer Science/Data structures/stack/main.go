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
	fmt.Println("stack len:", stk.Size())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Size())

	fmt.Println(stk.Pop())
	fmt.Println("stack len:", stk.Size())

	fmt.Println(stk.Pop())
	fmt.Println("stack len:", stk.Size())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())

	fmt.Println("stack len:", stk.Size())

	stk.Push(10)
	stk.Push(9)
	stk.Push(8)
	fmt.Println(stk.Size())

	fmt.Println(stk.PopAll())
	fmt.Println(stk.Size())

	// Fetch a singleton stack instance.
	stack := stack.Single[int]()

	// Use the stack.
	stack.Push(10)
	stack.Push(20)

	val, _ := stack.Pop()
	fmt.Println(val)

	// Release the stack back to the pool.
	stack.Release()

}
