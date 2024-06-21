package data_structures

// use a two-pointer approach: 1. insertPointer 2. uniquePointer
func removeDuplicates(nums []int) int {
	l, r := 0, 0
	for r < len(nums) {
		// duplicate
		if nums[l] == nums[r] {
			r++
		} else { // different so swap in
			nums[l+1] = nums[r]
			l++
			r++
		}
	}
	return l + 1
}
