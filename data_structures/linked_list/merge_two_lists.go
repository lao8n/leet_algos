package data_structures

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// choice: 1. for loop 2. recursion
	// choice: 1. in-place 2. new result
	// choice: 1. build from tail : more natural recursion but O(2n)
	//         2. build from head harder but single pass splicing in l2 values
	list1 = &ListNode{0, list1} // add to head of list so can use .Next test
	result := list1             // need head of list
	for list1.Next != nil {
		if list2 != nil {
			if list1.Next.Val > list2.Val {
				list1.Next = &ListNode{list2.Val, list1.Next} // splice in list2 val between list1 and l1.Next
				list2 = list2.Next                            // move onto next l2
			}
		}
		list1 = list1.Next
	}
	list1.Next = list2 // if list2 has any remaining elements add them on to the end
	return result.Next
}
