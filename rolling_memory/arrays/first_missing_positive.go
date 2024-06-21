package data_structures

// clarifying questions
// * only one missing number?

// approaches
// * brute force loop through looking for smallest number each time
// * heap of positive numbers where get minimum of heap to see if any missing - then pop off until missing value - heap is O(k logn) // where k is missing number
// * get max and min - then use euler's formula to figure out expected sum vs actual sum - doesn't work if duplicates
// * for constant space need to use more values in array trick

// specifics
// * can't use a seen array because that costs O(n) space
// * how handle duplicates?
func firstMissingPositive(nums []int) int {
	// set all 0 & neg numbers to 1
	seen1 := false
	for i, num := range nums {
		if num < 1 || num > len(nums) { // neg numbers & large numbers to 0
			nums[i] = 1
		}
		if num == 1 {
			seen1 = true
		}
	}
	if !seen1 {
		return 1
	}
	if len(nums) == 1 { // handle case where mapping to 1 index doesn't work
		return 2
	}

	// use index as neg to indicate number is seen
	for _, num := range nums {
		absNum := abs(num)
		// mark num as seen
		if absNum == len(nums) {
			nums[0] = -abs(nums[0])
		} else if nums[absNum] > 0 { // still unseen
			nums[absNum] *= -1 // make neg to mark as seen
		}
	}

	// loop through indices to see which are neg
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return len(nums)
	}
	return len(nums) + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
