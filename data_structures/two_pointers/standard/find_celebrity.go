package data_structures

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
// approach
// * loop through candidate celebrities until none? O(n)
// * loop through all pairs O(n/2 logn)
// specifics
// * everytime you ask a question you can eliminate one person

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		possCeleb, i, checkCount := 0, 1, 0
		// for every poss celeb need to check it with the next n
		for i < 2*n && checkCount < n {
			if i%n == possCeleb%n {
				checkCount++
			} else if knows(i%n, possCeleb%n) && !knows(possCeleb%n, i%n) {
				checkCount++
			} else {
				// i does not know poss celeb -> therefore need new candidate celeb
				possCeleb = i
				checkCount = 1
			}
			i++
		}
		if i == 2*n {
			return -1
		}
		return possCeleb % n
	}
}
