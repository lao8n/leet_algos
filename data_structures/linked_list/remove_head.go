package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// clarifying questions
// * 1-th numbered not 0-thed
// *
// approaches
// * 2 loops - 1st to find out how many in list 2nd to loop through to node to delete O(l + l - n)
// * recurse back - loop forward O(l) and loop backwards O(n) for O(l + n) -> which one preferred depends upon how large n is relative to l - n
// * could build a map of pointers to jump back to costs O(n) space but O(l) time
// * 2 pointer build new list from scratch - at n intervals
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var start ListNode
	start.Next = head
	_, node := nodeBeforeDeletion(&start, n)
	if node != nil && node.Next != nil {
		node.Next = node.Next.Next
	}
	return start.Next
}

func nodeBeforeDeletion(head *ListNode, n int) (int, *ListNode) {
	// base case
	if head == nil {
		return n, nil
	}
	// recursive case
	b, node := nodeBeforeDeletion(head.Next, n)
	if node != nil {
		return 0, node
	}
	if b == 0 {
		return 0, head
	}
	return b - 1, nil
}
