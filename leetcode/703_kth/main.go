package main

import "container/heap"

func main() {}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k      int
	stream *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)
	return KthLargest{k: k, stream: &h}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.stream, val)
	for this.stream.Len() > this.k {
		heap.Pop(this.stream)
	}
	return (*this.stream)[0]
}
