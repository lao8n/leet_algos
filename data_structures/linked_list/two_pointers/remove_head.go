package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// two pointers - one is to find the end and the other is n behind
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{-1, head}

	cur, prevOfRemoval := dummyHead, dummyHead

	for cur.Next != nil {

		// n step delay for prevOfRemoval
		if n <= 0 {
			prevOfRemoval = prevOfRemoval.Next
		}

		cur = cur.Next

		n -= 1
	}

	// Remove the N-th node from end of list
	nthNode := prevOfRemoval.Next
	prevOfRemoval.Next = nthNode.Next

	return dummyHead.Next
}
