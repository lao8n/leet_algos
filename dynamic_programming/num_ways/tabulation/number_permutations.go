package data_structures

// clarifying question
// * inversion = i < j && nums[i] > nums[j]
// * [end_i, cnt_i] = end index & inversion count of each requirement
// * num  of permutations

// TODO: don't understand this
func numberOfPermutations(n int, requirements [][]int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 401) // max count is 400
	}

	// requirements in a single array
	req := make([]int, n+1)
	for i := range req {
		req[i] = -1
	}
	for _, r := range requirements {
		e, c := r[0], r[1]
		req[e+1] = c
	}

	const mod = 1e9 + 7
	if req[1] <= 0 { // first requirement
		dp[1][0] = 1 // n = 1, count = 0
	}
	for i := 2; i <= n; i++ {
		for c := 0; c < i; c++ {
			for j := 0; j+c <= 400; j++ {
				if req[i] != -1 && j+c != req[i] {
					continue
				}
				dp[i][j+c] += dp[i-1][j]
				dp[i][j+c] %= mod
			}
		}
	}
	return dp[n][req[n]]
}
