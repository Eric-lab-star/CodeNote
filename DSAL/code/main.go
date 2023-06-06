package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := &iHeap{3, 6, 2}
	heap.Init(h)
	heap.Push(h, 8)
	min := heap.Pop(h)
	fmt.Println(min)

}

type iHeap []int

func (h iHeap) Len() int           { return len(h) }
func (h iHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h iHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *iHeap) Push(a interface{}) {
	*h = append(*h, a.(int))
}

func (h *iHeap) Pop() interface{} {
	n := len(*h)
	prev := *h
	x := (*h)[n-1]
	*h = prev[:n-1]
	return x
}
