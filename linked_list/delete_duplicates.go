package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// approaches
// * recursive with accumulator count vs iterative with count
// * in-place vs new list

// specifics
// * all duplicates
// * use dummy or sentinel head
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := ListNode{Val: 101, Next: head}
	cur, prev, next := &dummy, &dummy, dummy.Next
	count := 0 // val is 0 as default
	for next != nil {
		// fmt.Println(prev.Val, cur.Val, next.Val, count)
		// duplicate
		if cur.Val == next.Val {
			count++
			next = next.Next
			// don't advance prev as want to update pointer
		} else {
			if count == 1 { // ignore if count is 0
				// we want to keep this value
				prev.Next = cur
				prev = cur
			} else { // get rid of duplicates keep prev as is
				prev.Next = nil
			}
			count = 1
			cur = next
			next = next.Next
		}
	}
	// end
	if count == 1 {
		prev.Next = cur
	} else {
		prev.Next = nil
	}

	return dummy.Next
}
