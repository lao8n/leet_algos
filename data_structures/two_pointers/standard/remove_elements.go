package data_structures

// approaches
// * swap to end (or beginning) as order doesn't matter O(n)
func removeElement(nums []int, val int) int {
	swapIndex := len(nums) - 1
	i := 0
	for i <= swapIndex {
		if nums[i] == val {
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			// don't increment i as now need to process swap index value
			swapIndex--
		} else {
			i++
		}
	}
	return i
}
