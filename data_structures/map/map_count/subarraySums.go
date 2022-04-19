package data_structures

func subarraySum(nums []int, k int) int {
	// two-pointers doesn't work because of negative numbers
	// 1. naive try every start and stop index for O(n^3)
	// 2. cumulative sum array with start and end points O(n^2) with O(n) space
	// 3. end-point and sum at the same time O(n^2) O(1) space
	// 4. hash-map - key trick is that if the difference in cumulative sum
	// between two point is k then so is the sum between those two points.
	// store a map with sum -> freq of sum O(n) time O(n) space
	var count, sum int
	sumFreq := map[int]int{0: 1} // sum -> freq where need 0 : 1 for init
	for _, v := range nums {
		sum += v
		if freq, ok := sumFreq[sum-k]; ok {
			count += freq
		}
		sumFreq[sum]++
	}
	return count
}
