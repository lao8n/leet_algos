package data_structures

func minOperations(nums []int) int {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			if i+2 < len(nums) {
				for j := i; j < i+3; j++ {
					nums[j] = 1 - nums[j]
				}
				count++
			} else {
				return -1
			}
		}
		i++
	}
	return count
}
