// Heap
package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 9, 4}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}

}

type IntegerHeap []int

// implementing sort.Interface which is required for heap.Interface
func (iheap IntegerHeap) Len() int           { return len(iheap) }          // length of iheap
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] > iheap[j] } //ascending order
func (iheap IntegerHeap) Swap(i, j int)      { iheap[i], iheap[j] = iheap[j], iheap[i] }

// implementing heap.Interface
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	previous := *iheap
	n := len(previous)
	x1 := previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}
