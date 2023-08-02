package data_structures

// these are interconnected subproblems combined for a solution -> dynamic programming
// can either do bottom up tabulation or top-down recursion with memoization
func deleteAndEarn(nums []int) int {
	numsMap := make(map[int]int)
	maxNum := 0
	for _, n := range nums {
		numsMap[n] += n // store the number of points
		if n > maxNum {
			maxNum = n
		}
	}
	memoizedEarnMap := make(map[int]int)
	return max_points(maxNum, numsMap, memoizedEarnMap)
}

// big mistake i made was trying to memoize based on the set of keys in nums... much better is to
// memoize based upon some function like up to a certain amount, but very counter-intuitive to do it
// on all values up to a value i.e. max_points up to all values <= value
func max_points(num int, numsMap map[int]int, memoizedEarnMap map[int]int) int {
	// base cases
	if num == 0 {
		return 0
	} else if num == 1 {
		return numsMap[1]
	}
	// memoized values
	if earn, ok := memoizedEarnMap[num]; ok {
		return earn
	}
	// recursive cases - one of two options
	// choose the middle value
	middleValue := max_points(num-1, numsMap, memoizedEarnMap)
	// or get its neighbours
	neighbours := max_points(num-2, numsMap, memoizedEarnMap) + numsMap[num]
	memoizedEarnMap[num] = neighbours
	if middleValue > neighbours {
		memoizedEarnMap[num] = middleValue
		return middleValue
	}
	memoizedEarnMap[num] = neighbours
	return neighbours
}
