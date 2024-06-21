package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// approaches
// 1. iterate through linked list O(n) convert to list, sort O(n logn), then convert back to linked list O(n)
// 2. iterate through linked list and do an insertion sort O(n^2)
// 3. some variation of quick or merge sort i.e. recursively - how split up? need to know lengths.. merge sort is divide and conquer i.e. work is done on the merging
func sortList(head *ListNode) *ListNode {
	// find len - rather than finding len could use fast and slow pointer to find mid
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	return mergeSort(head, length)
}

func mergeSort(node *ListNode, length int) *ListNode {
	// base case
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}

	// break into two pieces left and right
	leftPartition, leftLen, rightPartition, rightLen := findMid(node, length)
	l := mergeSort(leftPartition, leftLen)
	r := mergeSort(rightPartition, rightLen)

	// merge them together
	dummy := &ListNode{}
	current := dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			current.Next = l
			l = l.Next
		} else {
			current.Next = r
			r = r.Next
		}
		current = current.Next
	}
	for l != nil {
		current.Next = l
		l = l.Next
		current = current.Next
	}
	for r != nil {
		current.Next = r
		r = r.Next
		current = current.Next
	}
	return dummy.Next
}

func findMid(node *ListNode, length int) (*ListNode, int, *ListNode, int) {
	leftLen := 0
	cur := node
	for leftLen = 1; leftLen < length/2; leftLen++ {
		cur = cur.Next
	}
	mid := cur.Next
	cur.Next = nil // stop this linked list

	return node, leftLen, mid, length - leftLen
}
