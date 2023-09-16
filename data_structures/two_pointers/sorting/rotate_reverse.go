package data_structures

// copy -> create new slice and append to old slice O(n) space but O(1) time
// but have to do it in place

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	// two pointers
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
