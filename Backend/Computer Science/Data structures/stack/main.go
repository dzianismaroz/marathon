package main

import (
	"fmt"

	stack "dzianismaroz.github.com/marathon/stack/pkg"
)

func main() {
	stk := stack.NewStack[int]()
	stk.Push(10)
	stk.Push(9)
	stk.Push(8)
	fmt.Println("stack len:", stk.Len())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Len())

	fmt.Println(stk.Pop())
	fmt.Println("stack len:", stk.Len())
}