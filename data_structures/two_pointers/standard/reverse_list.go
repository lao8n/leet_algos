package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// reverse linked list can either use an accumulator
// or can swap pointers
//
// x[1, y] -> y[2, z] -> z[3, nil]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		// swap
		cur.Next = prev

		// update
		prev = cur
		cur = next
	}
	return prev
}
