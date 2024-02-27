package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// clarifying questions
//
// approaches
// * create new list and add to end
// * repair even nubmers in place
func oddEvenList(head *ListNode) *ListNode {
	// loop over list
	cur := head
	oddDummy, evenDummy := new(ListNode), new(ListNode)
	lastOdd, lastEven := oddDummy, evenDummy
	odd := true
	for cur != nil {
		if odd {
			// if odd repair
			lastOdd.Next = cur
			lastOdd = lastOdd.Next
		} else {
			// if even add to even list
			lastEven.Next = cur
			lastEven = lastEven.Next
		}
		cur = cur.Next
		// flip
		odd = !odd
	}
	lastEven.Next = nil // fix cycle
	// attach together
	lastOdd.Next = evenDummy.Next
	return oddDummy.Next
}
