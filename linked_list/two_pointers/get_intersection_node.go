package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// if intersection then need to be of the same length
	aLen, bLen := getLength(headA), getLength(headB)
	// pick shorter and start at that index for longer
	for aLen > bLen {
		// a is longer
		headA = headA.Next
		aLen--
	}
	for bLen > aLen {
		// b is longer
		headB = headB.Next
		bLen--
	}
	// compare each node
	// if same then either candidate intersection or maintain previous candidate
	var intersectionNode *ListNode
	for headA != nil && headB != nil {
		if headA == headB {
			if intersectionNode == nil { // only update if nil
				intersectionNode = headA
			}
			// if different then need to start again
		} else {
			intersectionNode = nil
		}
		headA = headA.Next
		headB = headB.Next
	}
	return intersectionNode
}

func getLength(node *ListNode) int {
	if node == nil {
		return 0
	}
	return getLength(node.Next) + 1
}
