package data_structures

import (
	"container/heap"
)

// kth largest -> heap

// heap has two steps
// 1. heapification O(n)
// 2. removal of k elements which is O(k logn)

// two approaches to heap in golang
// 1. interface approach
// 2. array manipulation
func findKthLargest(nums []int, k int) int {
	h := make(Heap, len(nums))
	copy(h, nums)
	heap.Init(&h)
	for ; k > 1; k-- {
		heap.Pop(&h)
	}
	return h[0]
}

// heap implemenation
type Heap []int

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[0:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
