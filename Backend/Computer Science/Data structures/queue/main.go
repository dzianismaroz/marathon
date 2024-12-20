package main

import (
	"fmt"

	"github.com/dzianismaroz/marathon/queue/queue"
)

func main() {
	q := queue.New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println("Size:", q.Size())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println("Size:", q.Size())
	fmt.Println(q.Peek())

	fmt.Println("Size:", q.Size())
	fmt.Println(q.Pop())

	fmt.Println("Size:", q.Size())

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
