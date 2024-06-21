package data_structures

// two pointers : current num, move to
func moveZeroes(nums []int) {
	moveTo := 0
	for i, num := range nums {
		// we want 1 -> 0 and 0 to 1 where 1 is i and 0 is moveTo
		if i != moveTo {
			nums[i], nums[moveTo] = nums[moveTo], nums[i]
		}
		if num != 0 {
			moveTo++
		}
	}
}
