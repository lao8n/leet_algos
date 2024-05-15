package data_structures

// clarifying questions
// * only one missing number?

// approaches
// * brute force loop through looking for smallest number each time
// * heap of positive numbers where get minimum of heap to see if any missing - then pop off until missing value - heap is O(k logn) // where k is missing number
// * get max and min - then use euler's formula to figure out expected sum vs actual sum - doesn't work if duplicates
// * for constant space need to use more values in array trick
// * loop through numbers after a cycle sort

// specifics
// * can't use a seen array because that costs O(n) space
// * how handle duplicates?
func firstMissingPositive(nums []int) int {
	// cycle sort to place positive elements smaller than n
	for i := 0; i < len(nums); {
		correctIdx := nums[i] - 1
		// swap positive numbers
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[correctIdx] {
			nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
		} else {
			i++
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
