package data_structures

// this is a special case of lis
// lis can be solved with O(n logm) where m is the length of sequences
// intuitively you need a 2d array to store increasing sequences for all i to k
// but only the smallest elemnt for each ln needs to be tracked..
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// lis tabulation
	lis := []int{nums[0]} // for a sequence of length i what is its last element?
	for i := 1; i < len(nums); i++ {
		// longer increasing subsequence found
		if nums[i] > lis[len(lis)-1] {
			lis = append(lis, nums[i])
			if len(lis) >= 3 {
				return true
			}
		} else {
			// smaller number so do binary search to find the smallest
			j := binarySearch(lis, 0, len(lis)-1, nums[i]) // sort.Search(len(lis), func(k int) bool { return lis[k] >= nums[i] })
			lis[j] = nums[i]
		}
	}
	return false
}

// find smallest number that is greater >= cur
func binarySearch(lis []int, left int, right int, cur int) int {
	for left < right {
		mid := (left + right) / 2
		if lis[mid] >= cur { // [2, 4, 5] cur = 3 -> [2, 4]
			// search left
			right = mid
		} else if lis[mid] < cur { // [2, 4, 5] cur = 5 -> [4, 5]
			// search right
			left = mid + 1
		}
	}
	return left
}
