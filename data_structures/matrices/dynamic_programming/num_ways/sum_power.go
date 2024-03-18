package data_structures

func sumOfPower(nums []int, k int) int {
	n := len(nums)
	// setup dp
	dp := make([][]int, n+1) // (n + 1) x (k + 1) using numbers to n, how many ways to achieve sum up to k
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1

	// loop
	for _, num := range nums { // we add num
		for i := n; i > 0; i-- { // go backwards so you can previous rounds numbers
			for x := num; x <= k; x++ { // all x between num and k
				// all the ways to get x incl. i but now have new way involving not using i and getting x - num
				dp[i][x] = (dp[i][x] + dp[i-1][x-num]) % 1000000007
			}
		}
	}
	result := 0
	cur := 1
	for i := n; i > 0; i-- {
		result = (result + cur*dp[i][k]) % 1000000007
		cur = cur * 2 % 1000000007 // double current for all subseqs with and without i
	}
	return result
}
