package main

import "container/heap"

func main() {}

type (
	StonesHeap []int
)

func (h StonesHeap) Len() int           { return len(h) }
func (h StonesHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h StonesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *StonesHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *StonesHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	sHeap := make(StonesHeap, 0, len(stones))
	copy(sHeap, stones)

	h := StonesHeap(sHeap)
	heap.Init(&h)

	for _, stone := range stones {
		heap.Push(&h, stone)
	}

	for h.Len() > 1 {
		first := heap.Pop(&h).(int)
		second := heap.Pop(&h).(int)

		if first != second {
			heap.Push(&h, first-second)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(&h).(int)
}
