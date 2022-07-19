package data_structures

import "math"

// approach 1 : naive
// 1. for loop through all starting positions 2. for loop through all ending positions checking if palindrome = O(n^2)
// improvements : 1. memoize calculations

// approach 2: two pointers either in then out or out then in...
// 1. for loop through possible mid-points 2. go left and right until not palindrome 3. can we memoize anything? = O(n x k)
// example
// if have palindrome and want to see if there is a new palindrome the only possibilities are:
// 1. new palindrome starts at end [..., 3, 4, 1, 4, 3]
// 2. new palindrom includes all of sequence with midpoint after end

// approach 3: 1. build map with value and locations 2. two-pointers on starting and end locations but have to try every combination?
// = O(n) +
func longestPalindrome(s string) string {
	midpoint, left, right, longestLeft, longestRight := 0.0, 0, 0, 0, 0
	for int(midpoint) < len(s) {
		left, right = int(math.Floor(midpoint)), int(math.Ceil(midpoint))
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left > longestRight-longestLeft {
				longestLeft, longestRight = left, right
			}
			left--
			right++
		}
		midpoint = midpoint + 0.5
	}
	return s[longestLeft : longestRight+1]
}
