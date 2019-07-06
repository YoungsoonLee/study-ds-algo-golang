package main

import (
	"container/heap"
	"fmt"
)

// integerheap a type
type IntegreHeap []int

func (h IntegreHeap) Len() int { return len(h) }

func (h IntegreHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntegreHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntegreHeap) Push(heapintf interface{}) {
	*h = append(*h, heapintf.(int))
}

func (h *IntegreHeap) Pop() interface{} {
	var n int
	var x1 int
	var prev IntegreHeap = *h
	n = len(prev)
	x1 = prev[n-1]
	*h = prev[0 : n-1]
	return x1
}

func main() {
	var ih *IntegreHeap = &IntegreHeap{1, 4, 5}

	heap.Init(ih)
	heap.Push(ih, 2)
	fmt.Printf("minimum: %d\n", (*ih)[0])
	for ih.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(ih))
	}
}
