package data_structures

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// for loop through [x] -> [y] -> [z]
// ListNode{x, nil}
// result = {x}
// head = y
// ListNode{y, x}
// result = {y}
// head = z
// ListNode{z, y}
func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		temp := ListNode{Val: head.Val, Next: result}
		result = &temp
		head = head.Next
	}
	return result
}
