package data_structures

// interconnected problems combined for a solution -> dynamic programming esp as num of ways

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// given remaining target easier to manage with recursion
// 3 components of dp
// 1. base cases
// 2. recurrence -> add one of 1-6 to target
// 3. state -> remaining number of dice rolls, remaining target

// we don't care about duplicates
// seems like top down is too many things to calculate - even if memoized - so let's try bottom up

// time complexity O(target * n * k)
// space complexity O(target * n)
func numRollsToTarget(n int, k int, target int) int {
	tabulation := make([][]int, target+1)
	for i, _ := range tabulation {
		tabulation[i] = make([]int, n+1)
	}

	for i := 1; i <= target; i++ {
		for j := 1; j <= n; j++ {
			// num ways that can get target i with j throws
			numWays := 0
			if j == 1 {
				if i <= k {
					numWays = 1
				}
			} else {
				for c := 1; c <= k; c++ {
					if i-c >= 0 {
						if i-c >= 0 && j-1 >= 1 {
							numWays += tabulation[i-c][j-1]
						}
					}
				}
			}
			tabulation[i][j] = numWays % 1000000007
		}
	}
	return tabulation[target][n]
}
