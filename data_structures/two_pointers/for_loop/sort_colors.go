package data_structures

// order: red -> white -> blue
// constraints = no sort function, & in-place

// naive 3x O(n) passes, when find colour swap with last known index
// single pass = if red -> left index, if blue -> right index, white leave i.e. two-pointer approach
func sortColors(nums []int) {
	redIndex, blueIndex := 0, len(nums)-1
	for i := 0; i < len(nums); {
		n := nums[i]
		if n == 0 && i >= redIndex {
			swap(nums, i, redIndex)
			redIndex++
			i++
		} else if n == 2 && i <= blueIndex {
			swap(nums, i, blueIndex)
			blueIndex--
		} else {
			i++
		}
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
