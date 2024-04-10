package data_structures

// big clue is numbers only between 1 and n for cyclic sort

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		correctIdx := nums[i] - 1
		if nums[correctIdx] != nums[i] {
			// swap
			nums[correctIdx], nums[i] = nums[i], nums[correctIdx]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return nums[i]
		}
	}
	return 0
}
