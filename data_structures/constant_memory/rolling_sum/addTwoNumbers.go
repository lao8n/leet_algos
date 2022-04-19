package data_structures

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// reverse-order = natural way to get number size is build from smallest digit
	// but would require either O(n) storage and O(2n) pass
	// single-pass = need to know power of 10 to multiply by so either 1. use len but requires O(n) or 2. O(1) storage and multiply on fly
	// building digits = single pass is intuitive, extract digits with / 10 remainder
	// choices: 1. recursive 2. for loop
	// choices: 1. calculate sum then create output 2. calculate sum and output at the same time
	cur := new(ListNode)
	ret := cur
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			sum % 10,
			nil,
		}
		cur = cur.Next
		sum = sum / 10
	}

	if sum > 0 {
		cur.Next = &ListNode{sum, nil}
	}

	return ret.Next
}
