package data_structures

// clarifying questions
// * can include negatives numbers? yes
// * can includes duplicate values? just not duplicate indices
// * care about ordering of triplets? no
// approaches
// * for loop over two sum by map O(n x n) time but O(n) space
// * sort and then use 3 pointers O(n logn) + O(n x n) where 2nd n is the for the two pointers
func threeSum(nums []int) [][]int {
	solution := make([][]int, 0)
	for i, num1 := range nums {
		complementMap := make(map[int]bool, len(nums))
		target := -num1 // we want num + other two = 0 i.e. we want other two = -num
		// single pass to find complement
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			// check if complement exists - then have solution
			if _, ok := complementMap[target-num2]; ok {
				solution = append(solution, []int{num1, target - num2, num2})
			}
			// otherwise add to map and proceed
			complementMap[num2] = true
		}
	}
	return solution
}
