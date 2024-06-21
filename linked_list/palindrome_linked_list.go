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
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	// check palindrome
	l, r := 0, len(values)-1
	for l < r {
		if values[l] != values[r] {
			return false
		}
		l++
		r--
	}
	return true
}
