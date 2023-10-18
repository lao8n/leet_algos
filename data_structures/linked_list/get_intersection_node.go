package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// clarifying questions?
// * intersection based upon node pointer or value? pointer because 1 example

// approach?
// * brute force is to for every head A check to see if same as every head b but this is O(n^2)
// * cannot recurse down in parallel because could be different lengths - recursing backwards is difficult
// * use a data structure such as a list to get reverse order
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aStack := make([]*ListNode, 0)
	bStack := make([]*ListNode, 0)

	// add to stacks all nodes
	for headA != nil {
		aStack = append(aStack, headA)
		headA = headA.Next
	}
	for headB != nil {
		bStack = append(bStack, headB)
		headB = headB.Next
	}
	// pop from both stacks (LIFO) untile don't match
	aI, bI := len(aStack)-1, len(bStack)-1
	var intersectionNode *ListNode
	for aI >= 0 && bI >= 0 {
		if aStack[aI] == bStack[bI] {
			intersectionNode = aStack[aI]
		}
		aI--
		bI--
	}
	return intersectionNode
}
