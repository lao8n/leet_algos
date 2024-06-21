// clarifying questions
// * has to be a divisor - can be 1? but cannot be 0 or n

// approaches
// * dynamic programming where build up from 1 to n for available moves - choosing between.. not sure how to quickly come up with all the valid divisors? do i just have to cycle through?
// * is there some quick rule? yes there is n % 2 == 0 - if given an odd number always have to return an odd number

package data_structures

func divisorGame(n int) bool {
	// dp table
	dp := make([]bool, n+1) // so can have numbers without - 1, default false

	// base cases
	dp[1] = false // 1 remaining alice loses

	for i := 2; i <= n; i++ {
		// look to see if any possibleX are false -> then can win
		for _, x := range possibleX(i) {
			if !dp[i-x] {
				dp[i] = true
			}
		}
	}
	return dp[n]
}

func possibleX(n int) []int {
	possible := make([]int, 0)
	for x := 1; x < n; x++ {
		if n%x == 0 {
			possible = append(possible, x)
		}
	}
	return possible
}
