package data_structures

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// clarifying questions
// approaches
// * get entire list O(n) then check for palindrome with two pointers for O(n) again
// * single-pass - assume each point is possible middle and have rolling true or false - might need to have, problem is there isn't really overlap in work done [5, 3, 1, 3] -> if stop here not valid palindrome - could determine if palindrome from middle but first we hypothesis 3 say 5 and 1 so cannot be then we hypothesise between 3 & 1 -> no work saved
func isPalindrome(head *ListNode) bool {
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverse(firstHalfEnd.Next)

	// check whether palindrome
	palindrome := true
	f, s := head, secondHalfStart
	for palindrome && s != nil {
		if f.Val != s.Val {
			palindrome = false
		}
		f = f.Next
		s = s.Next
	}
	// restore list
	firstHalfEnd.Next = reverse(secondHalfStart)

	return palindrome
}

func endOfFirstHalf(node *ListNode) *ListNode {
	fast, slow := node, node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}
