package data_structures

import "container/heap"

// effectively we need to maintain largest stones with deletion and insert -> heap
// two steps
// 1. heapification O(n)
// 2. remove min O(1) 2x + insert O(logn)
func lastStoneWeight(stones []int) int {
	h := make(Heap, len(stones))
	copy(h, stones)
	heap.Init(&h)
	for h.Len() > 1 {
		largest, secondLargest := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if largest != secondLargest {
			heap.Push(&h, largest-secondLargest)
		}
	}
	if h.Len() != 0 {
		return heap.Pop(&h).(int)
	}
	return 0
}

type Heap []int

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] } // we want max heap
