package data_structures

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// choice 1. look at next or look at current

func reverseList(head *ListNode) *ListNode {
	return reverseListRec(head, nil)
}

func reverseListRec(head *ListNode, reversed *ListNode) *ListNode {
	// base case
	if head == nil {
		return reversed
	}
	// recursive case
	next := ListNode{Val: head.Val, Next: reversed}
	result := reverseListRec(head.Next, &next)
	return result
}
