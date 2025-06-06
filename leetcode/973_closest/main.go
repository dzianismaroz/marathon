package main

import (
	"container/heap"
	"math"
)

func main() {}

type (
	Point struct {
		distance float64
		point    []int
	}

	Closest []Point
)

func (c Closest) Len() int           { return len(c) }
func (c Closest) Less(i, j int) bool { return c[i].distance < c[j].distance }
func (c Closest) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *Closest) Push(point any) {
	*c = append(*c, point.(Point))
}

func (c *Closest) Pop() any {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := Closest{}
	heap.Init(&h)

	for _, p := range points {
		heap.Push(&h, Point{math.Hypot(float64(p[0]), float64(p[1])), p})
	}

	result := make([][]int, k)

	for i := range k {
		result[i] = heap.Pop(&h).(Point).point
	}

	return result
}
