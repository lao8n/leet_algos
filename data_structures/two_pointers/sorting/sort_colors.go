package data_structures

// order: red -> white -> blue
// constraints = no sort function, & in-place

// naive 3x O(n) passes, when find colour swap with last known index
// single pass = if red -> left index, if blue -> right index, white leave i.e. two-pointer approach
func sortColors(nums []int) {
	redPointer, bluePointer := 0, len(nums)-1
	for i := 0; i < len(nums); {
		if nums[i] == 0 && i >= redPointer {
			nums[i], nums[redPointer] = nums[redPointer], nums[i]
			redPointer++
			i++
		} else if nums[i] == 2 && i <= bluePointer {
			nums[i], nums[bluePointer] = nums[bluePointer], nums[i]
			bluePointer--
		} else {
			i++
		}
	}
}
