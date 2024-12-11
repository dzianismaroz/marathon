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
	fmt.Print(stk.Size())
}
