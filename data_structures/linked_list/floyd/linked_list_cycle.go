package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// clarifying questions
// * is val unique or is it hte node pointer?

// approaches
// * could have a set of list nodes that have seen
// * rabbit & hare -> if rabbit goes at 2x the speed is it guaranteed they will end up in same point? if cycle has even number of nodes then [0, 1, 2, 3, 4, 5, 6] they will land on same node when 2x = x, but how does this work for odd number in cycle?
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false // reached end therefore must not be cycle
}
