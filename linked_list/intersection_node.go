package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		// traverse exclusive a and then c and then b
		if pA != nil {
			pA = pA.Next // find the length of a and c
		} else {
			pA = headB
		}
		// traverse exclusive b and then c and then a
		if pB != nil {
			pB = pB.Next // find the length of b and c
		} else {
			pB = headA
		}
	}
	return pA
}
