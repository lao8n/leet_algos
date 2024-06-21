package data_structures

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head
	groupStartPrev := dummy
	for head != nil {
		groupEnd := head
		for i := 1; i < k && groupEnd != nil; i++ {
			groupEnd = groupEnd.Next
		}
		if groupEnd == nil {
			break
		}
		nextGroupStart := groupEnd.Next
		groupEnd.Next = nil
		reversed, tail := reverseLinkedList(head, k)
		groupStartPrev.Next = reversed
		tail.Next = nextGroupStart
		groupStartPrev = tail
		head = nextGroupStart
	}
	return dummy.Next
}

func reverseLinkedList(head *ListNode, k int) (*ListNode, *ListNode) {
	var prev, tail *ListNode
	tail = head
	current := head
	for k > 0 {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
		k--
	}
	return prev, tail
}
