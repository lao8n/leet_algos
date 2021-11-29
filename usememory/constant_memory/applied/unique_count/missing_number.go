func missingNumber(nums []int) int {
	// sort number O(n logn) -> increment until find missing O(n)
	// sum all elements O(n) keep track of max and min and then use (max * min) / 2
	// but min = 0 so just do len * len + 1 / 2
	var sum int
	for _, v := range nums {
		sum += v
	}
	expectedSum := len(nums) * (len(nums) + 1) / 2
	return expectedSum - sum
}