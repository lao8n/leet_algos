package data_structures

import "container/heap"

// approaches
// * naive approach is to loop over window everytime
// * want to be able to 1. pop first value FIFO 2. know maximum immediately -> use combination of a linked list which is O(1) removal - or could use a fixed size list with modulus and starting index
// * keep track of max -> could use heap -> but how remove specific index value -> queue - with max
// [1  3  -1] -3  5  3  6  7       3 - store [val = 1, max = 3, count = 1]
// 1 [3  -1  -3] 5  3  6  7       3 - store [val = 3, max = 3, count = 1]
// 1  3 [-1  -3  5] 3  6  7       5 - store [val = -1, max = 5, count = 1]
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// * stack doesn't work because if had [1, 3, -1] then could have all neg numbers and -1 will become a max
// * use a heap where we have two criteria 1. max number 2. earliest index then we can push and pop correclty

// clarifying questions/specifics?
// * can k > nums?
// * decreasing numbers 5, 4, 3, 2, 1, 0, -1 - need every value stored
func maxSlidingWindow(nums []int, k int) []int {
	// base case
	if len(nums) == 0 {
		return nil
	}

	// initialize heap
	h := make(Heap, 0)
	heap.Init(&h)
	solution := make([]int, 0)

	// loop through k
	i := 0
	for i < k-1 && i < len(nums) {
		heap.Push(&h, data{val: nums[i], ind: i})
		i++
	}
	// loop through nums
	for i < len(nums) {
		heap.Push(&h, data{val: nums[i], ind: i})
		i++
		max := h.Peek()
		// update heap window if max is at last index
		for max.ind < i-k { // [1, -1] i = 0, k = 1 -1
			// pop all max until we are within our new window
			heap.Pop(&h)
			max = h.Peek()
		}
		solution = append(solution, max.val)
	}
	return solution
}

type data struct {
	val int
	ind int
}

type Heap []data

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(data))
}
func (h *Heap) Pop() interface{} {
	previous, popped := *h, data{}
	*h, popped = previous[:h.Len()-1], previous[h.Len()-1]
	return popped
}
func (h Heap) Peek() data {
	return h[0] // not h.Len() - 1
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(x, y int) bool { // want a max heap
	if h[x].val == h[y].val {
		return h[x].ind < h[y].ind
	}
	return h[x].val > h[y].val
}
func (h Heap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
