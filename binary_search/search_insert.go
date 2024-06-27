package data_structures

// approach
// * logn means we should use binary search - could do recursive or loop
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if target > nums[mid] {
		return mid + 1
	}
	return mid
}
