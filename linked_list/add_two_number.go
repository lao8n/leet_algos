package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// approaches
// * reverse inputs and carry
// * something clever with length of pointers like l1 + l2, find lengths of each with first pass
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Stack, l2Stack := make([]int, 0), make([]int, 0)
	for l1 != nil {
		l1Stack = append(l1Stack, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		l2Stack = append(l2Stack, l2.Val)
		l2 = l2.Next
	}
	// build from stacks
	i, j := len(l1Stack)-1, len(l2Stack)-1
	carry := 0
	var prev *ListNode
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += l1Stack[i]
		}
		if j >= 0 {
			sum += l2Stack[j]
		}
		digit := sum % 10
		carry = sum / 10
		prev = &ListNode{Val: digit, Next: prev}
		i--
		j--
	}
	if carry != 0 {
		return &ListNode{Val: carry, Next: prev}
	}
	return prev
}
