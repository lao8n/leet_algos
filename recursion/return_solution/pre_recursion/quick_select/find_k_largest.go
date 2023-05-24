package recursion

import "math/rand"

// kth largest -> sort
// with divide & conquer or conquer then divide -> latter = quick select
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left, right, kLargest int) int {
	// base cases
	if left == right {
		return nums[left]
	}

	// recursive = conquer then divide
	pivotIndex := left + rand.Intn(right-left) // necessary for sorted data
	numLeftOfPivot := partition(nums, left, right, pivotIndex)

	if kLargest == numLeftOfPivot {
		return nums[kLargest]
	}
	if kLargest < numLeftOfPivot {
		return quickSelect(nums, left, numLeftOfPivot-1, kLargest)
	}
	return quickSelect(nums, numLeftOfPivot+1, right, kLargest)
}

func partition(nums []int, left, right, pivotIndex int) int {
	pivot := nums[pivotIndex]
	// 1. pivot to end
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	// 2. all elements to left
	swapIndex := left
	for i := left; i <= right; i++ {
		if nums[i] > pivot {
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			swapIndex++
		}
	}
	// 3. pivot to final place
	nums[swapIndex], nums[right] = nums[right], nums[swapIndex]

	return swapIndex
}
