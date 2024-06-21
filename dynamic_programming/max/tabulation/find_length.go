package data_structures

// interconnected problems combined for a solution -> dynamic programming

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components
// 1. base cases
// 2. state -> index in nums1, index in nums2 - longest array starting at n1, n2 or longest array from n1, n2?
// 3. recurrence ->

// time complexity O(len(nums1) x len(nums2))
// space complexity O(len(nums1) x len(nums2))
func findLength(nums1 []int, nums2 []int) int {
	tabulation := make([][]int, len(nums1))
	max := 0
	for n1 := 0; n1 < len(nums1); n1++ {
		tabulation[n1] = make([]int, len(nums2))
		for n2 := 0; n2 < len(nums2); n2++ {
			if nums1[n1] == nums2[n2] {
				tabulation[n1][n2] += 1
				if n1 > 0 && n2 > 0 {
					tabulation[n1][n2] += tabulation[n1-1][n2-1]
				}
			}
			if tabulation[n1][n2] > max {
				max = tabulation[n1][n2]
			}
		}
	}
	return max
}
