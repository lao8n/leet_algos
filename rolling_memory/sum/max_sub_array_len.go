package data_structures

// clarifications
// * continuous?

// approaches
// * brute force try every start index and every end index O(n^3) as need to loop over array
// * rolling subarray - look up complement i.e. if rolling sum is 10, and target is 11, need -1 to be rolling sum?

// specifics
// * allowed to have negative numbers so cannot bail out early if exceed k
// * max length if same complement exists keep the first one
// * need to try all options
func maxSubArrayLen(nums []int, k int) int {
	sumMap := make(map[int]int) // rolling sum -> index (keep first)
	sum := 0
	maxLen := 0
	sumMap[0] = -1
	for i, num := range nums {
		// fmt.Println(sumMap, num, num + sum, k - num - sum)
		sum += num
		// check sum map
		// we want sum - prev = k - sum
		if idx, ok := sumMap[-(k - sum)]; ok {
			curLen := i - idx // -2 - which is before
			if curLen > maxLen {
				maxLen = curLen
			}
		}
		// add to sum map
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return maxLen
}
