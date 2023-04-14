packages data_structures
// interdependent subproblems combined for a solution -> dynamic programming
// dynamic programming has 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// max sum subarray problem -> kadane's algorithm (optimised form of 2 with O(1) space complexity)
// approaches
// 1. do kadane on array O(n) + kadane on beginning of array again
// 2. for middle segment either we take middle section or we take outer sections

func maxSubarraySumCircular(nums []int) int {
    maxSum := math.MinInt
    minSum := math.MaxInt
    rollingSum, rollingMaxSum, rollingMinSum := 0, 0, 0
    for i := 0; i < len(nums); i++ {
        // max sum logic
        if rollingMaxSum + nums[i] < nums[i] {
            rollingMaxSum = nums[i]
        } else {
            rollingMaxSum += nums[i]
        }
        if rollingMaxSum > maxSum {
            maxSum = rollingMaxSum
        }
        // min sum logic
        if rollingMinSum + nums[i] > nums[i] {
            rollingMinSum = nums[i]
        } else {
            rollingMinSum += nums[i]
        }
        if rollingMinSum < minSum {
            minSum = rollingMinSum
        }
        // sum
        rollingSum += nums[i]
        fmt.Println(minSum, maxSum, rollingSum)
    }
    if rollingSum != minSum && rollingSum - minSum > maxSum {
        return rollingSum - minSum
    }
    return maxSum
}