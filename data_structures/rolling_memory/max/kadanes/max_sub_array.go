package data_structures

// naive O(n2) = grid with all starting positions and all ending positions
// memoization = however for a given sI and eI, sum[eI - sI] = sum[eI] - sum[sI]
// O(n) = rolling sum -> then find max value and min value -> and that is largest sum
// problem that rolling max needs to follow rolling min
func maxSubArray(nums []int) int {
	currentRollingSum, maxRollingSum := nums[0], nums[0]
	// this is just kadane's algorithm
	for i := 1; i < len(nums); i++ {
		if currentRollingSum < 0 {
			currentRollingSum = 0
		}
		currentRollingSum += nums[i]
		if currentRollingSum > maxRollingSum {
			maxRollingSum = currentRollingSum
		}
	}
	return maxRollingSum
}
