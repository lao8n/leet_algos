package data_structures

import "container/heap"

// clarifying questions
// when we remove elements we remove evne if not less than k - just take two smalelst?

// approaches
// * calculation seems way too complicated to have a trick to it - answer is calculate manually with a min heap
func minOperations(nums []int, k int) int {
	// setup heap
	m := make(minHeap, len(nums))
	copy(m, nums)
	heap.Init(&m) // O(n)
	count := 0

	// loop through heap until condition met
	for m.Peek() < k {
		count++
		x := heap.Pop(&m).(int)
		y := heap.Pop(&m).(int)
		push := min(x, y)*2 + max(x, y)
		heap.Push(&m, push)
	}
	return count
}

type minHeap []int

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *minHeap) Pop() interface{} {
	popped, orig, n := 0, *m, m.Len()
	*m, popped = orig[:n-1], orig[n-1]
	return popped
}
func (m minHeap) Peek() int {
	return m[0]
}
func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Swap(x, y int) {
	m[x], m[y] = m[y], m[x]
}
func (m minHeap) Less(x, y int) bool {
	return m[x] < m[y]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
