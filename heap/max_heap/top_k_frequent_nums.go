package data_structures

import "container/heap"

// kth largest -> heap

// heap has two steps
// 1. heapification O(n)
// 2. removal of k elements which is O(k log n)

// two approaches to heap in golang
// 1. interface approach
// 2. array manipulation

// time complexity = O(k logn)
// space complexity = O(n)
func topKFrequent(nums []int, k int) []int {
	// O(n) get frequencies
	hMap := make(map[int]int, len(nums))
	for _, num := range nums {
		hMap[num]++
	}
	// O(n) heapify
	h := make(Heap, len(hMap))
	i := 0
	for k, v := range hMap {
		h[i] = freq{num: k, freq: v}
		i++
	}
	heap.Init(&h)

	// Pop heap
	popped := make([]int, k)
	for i := 0; i < k; i++ {
		popped[i] = heap.Pop(&h).(freq).num
	}
	return popped
}

type freq struct {
	num  int
	freq int
}
type Heap []freq

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(freq)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), freq{}
	*h, popped = previous[0:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].freq > h[j].freq } // make it greater than
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
