package data_structures

// naive approach: just loop through all numbers to check if existence O(n)
// strictly ascending: loop through until reach number that is too larger
// rotated array: either go forwards from 0th index or backwards from n-1 index until find O(n/2)
// binary search: partition in half, and pick side
func search(nums []int, target int) int {
	return partition(nums, target, 0, len(nums)-1)
}

// can simplify by thinking about two ends and mid-point explicitly.
func partition(nums []int, target int, i1 int, i2 int) int {
	// base cases
	if i2-i1 <= 1 && nums[i1] != target && nums[i2] != target {
		return -1
	} else if target == nums[i1] {
		return i1
	} else if target == nums[i2] {
		return i2
	}
	i2 = i1 + (i2-i1)/2
	v1, v2 := nums[i1], nums[i2]
	// compare to target
	if v1 <= target && target <= v2 {
		return partition(nums, target, i1, i2)
		// nums = [4,5,6,7,0,1,2], v1 = 4, v2 = 0, target = 6
	} else if v2 <= target && target <= v1 {
		return partition(nums, target, i2, len(nums)-1)
		// nums = [4,5,6,7,0,1,2], v1 = 4, v2 = 7, target = 2
	} else if v1 < v2 {
		return partition(nums, target, i2, len(nums)-1)
		// nums = [4,5,6,7,0,1,2], v1 = 4, v2 = 0, target = 2
	} else { // v2 >= v1
		return partition(nums, target, i1, i2)
	}
}
