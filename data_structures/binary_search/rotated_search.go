package data_structures

// naive approach: just loop through all numbers to check if existence O(n)
// strictly ascending: loop through until reach number that is too larger
// rotated array: either go forwards from 0th index or backwards from n-1 index until find O(n/2)
// binary search: partition in half, and pick side
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // nums[mid] < nums[start]]
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
