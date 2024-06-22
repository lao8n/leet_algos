package data_structures

func minOperations(nums []int) int {
	count := 0
	i := 0
	for i < len(nums) {
		if count%2 == 1 {
			nums[i] ^= 1 // flip
		}
		if nums[i] == 0 {
			count++
		}
		i++
	}
	return count
}
