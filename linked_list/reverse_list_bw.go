package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// approaches
// * build new list in forward pass with accumulate
// * build new list on backward pass
func reverseList(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// recursive case
	reversed := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversed
}
