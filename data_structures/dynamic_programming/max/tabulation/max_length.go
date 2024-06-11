package data_structures

// clarifying questions
// * good = at most k indices i in the range [0, seq.length - 2] - i.e. number of different pairs
// approaches
// * two pointers - count number of different pairs
// * now understand can also remove elements in the middle - so two pointers works
// * not just about deleting once could delete multiple elements
// * max len is usually dp - what should the recurrence relation be. could probably get away with a 2d-array - but
// might as well have 1-d
// * can i work out from k = 0 seq ending at i seq ending at i + 1?
// * everytime you add a new element - you compare it to all previous elements and whether it adds a k and max subsequence length
// specifics
// * dp[i][j] represents the longest streak up to index i but not necessarily including index i and j
func maximumLength(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// setup dp
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	res := 1
	for j := 0; j <= k; j++ {
		bestLenForJ := 1
		numMap := make(map[int]int) // rather than for loop going back
		numMap[nums[0]] = 0
		for i := 1; i < n; i++ {
			dp[i][j] = 1
			// previous num -> add a len and add a k
			if i > 0 && j > 0 {
				bestLenForJ = max(bestLenForJ, dp[i-1][j-1]+1)
			}
			dp[i][j] = max(dp[i][j], bestLenForJ)
			// unless same num exists and then just add a len but same k
			if idx, ok := numMap[nums[i]]; ok {
				dp[i][j] = max(dp[i][j], dp[idx][j]+1)
			}
			numMap[nums[i]] = i
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
