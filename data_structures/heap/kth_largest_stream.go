package data_structures

import "container/heap"

// kth largest -> heap
// two steps
// 1. heapification O(n) - although in our case it is O(n logk) for insert
// 2. remove min O(1)

// what we want:
// a min heap of size k, where if the heap is full we only insert if it is bigger than the minimum
type KthLargest struct {
	k    int
	heap Heap
}

func Constructor(k int, nums []int) KthLargest {
	// Heapify O(n)
	h := make(Heap, len(nums))
	copy(h, nums)
	heap.Init(&h)
	return KthLargest{k: k, heap: h}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.heap, val)
	for this.heap.Len() > this.k {
		heap.Pop(&this.heap)
	}
	return this.heap.Peek()
}

type Heap []int

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h *Heap) Peek() int {
	return (*h)[0]
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
