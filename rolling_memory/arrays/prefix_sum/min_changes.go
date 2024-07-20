package data_structures

import "math"

// clarifying questions
// * n even
// * replace any element with 0-k
// * requirement = a[i] - a[n-i-1] // symmetric around middle i.e. difference between two values is X
// * all nums[i] < k?

// approaches
// * could loop through all 0..k && loop through nums to calculate changes needed but inefficient O(nk)
// * do range additions instead do O(n)

func minChanges(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, k+1)

	for i := 0; i < n/2; i++ {
		left, right := nums[i], nums[n-1-i]
		diff := abs(left - right)
		maxDiff := max(max(left, right), max(k-left, k-right)) // left or right to 0 or left or right to k
		prefixSum[0] += 1                                      // make every pair have 0 difference
		prefixSum[diff] -= 1                                   // no operation required for this sum, so offset prefixSum[0]
		if diff+1 <= k {
			prefixSum[diff+1] += 1
		}
		if maxDiff+1 <= k {
			prefixSum[maxDiff+1] += 1
		}
	}

	rangeSum := 0
	minSum := math.MaxInt
	for i := 0; i <= k; i++ {
		rangeSum += prefixSum[i]
		minSum = min(minSum, rangeSum)
	}
	return minSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
