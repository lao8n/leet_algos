package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// we are not given the previous node so we cannot update the pointer
// therefore we need the current node that is connected to it to stay connected..
// choice between looping through the entire end of list updating values?
// or just update a few
// [5, a] -> a[1, b] -> b[9, c] -> c[4, nil]
// [5, a] -> a[9, c] ->
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
