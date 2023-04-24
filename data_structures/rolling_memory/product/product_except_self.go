package data_structures

func productExceptSelf(nums []int) []int {
	// O(2n) space is a left and right matrix
	// can do better by using result as left matrix and calculating right on the fly with O(1)
	result := make([]int, len(nums))
	result[0] = 1 // ith index is product of all values < i
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	// update result with R index
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * r
		r *= nums[i]
	}
	return result
}
