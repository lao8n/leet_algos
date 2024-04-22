package data_structures

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// approaches
// * new list or in place - former is simpler and not that much more expensive
// * have a bunch of pointers and increment the lowest one -> could just loop through all the lists but there are up to 500 - maybe use a heap to get the lowest number?
// * if pointer per list approach that is O(kn) but also have to loop through k every round so that is O(k^2 n)
// * if use a heap for all numbers and just pop then you have a heap of size nk, to pop each value takes log nk (init is not relevant) therefore cost is O(nk lognk)
// * if we use a heap of just k elements then every element need to be pushed and popped for log k costs which is O(nk logk)
// * could have use divide and conquer for O(n logk) and O(n logk) time and space
// specifics
// * do we need to know what list? actually just need a pointer to start of each list?
func mergeKLists(lists []*ListNode) *ListNode {
	// create heap of k elements
	h := make(Heap, 0, len(lists))
	for _, list := range lists {
		if list != nil {
			h = append(h, list)
		}
	}
	heap.Init(&h)

	// create new list node
	dummy := ListNode{}
	prev := &dummy
	// loop until heap empty
	for h.Len() > 0 {
		// pop smallest value
		popped := heap.Pop(&h).(*ListNode)
		cur := &ListNode{Val: popped.Val}
		prev.Next = cur
		prev = cur

		// get next value and push to heap
		if popped.Next != nil {
			heap.Push(&h, popped.Next)
		}
	}
	// return list node
	return dummy.Next
}

type Heap []*ListNode

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *Heap) Pop() interface{} {
	popped := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return popped
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h Heap) Less(x, y int) bool {
	return h[x].Val < h[y].Val // want min heap
}
