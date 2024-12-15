package main

import (
	"fmt"

	"github.com/dzianismaroz/marathon/tree/tree"
)

func main() {
	t := tree.New[int]()

	t.Add(44)
	t.Add(18)
	t.Add(1)
	t.Add(2)
	t.Add(10)
	t.Add(8)

	fmt.Println("size: ", t.Size())

	fmt.Println(t.SortedDesc())
	fmt.Println(t.SortedAsc())

}
