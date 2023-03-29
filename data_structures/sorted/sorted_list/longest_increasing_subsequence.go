package data_structures

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches ot dynamic programming
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// however we can do better with a greedy approach where we maintain the best sequence
// so far (technically the final solution might not be a valid seq but we only care
// about the length)
// this isn't usually possible because the earlier choice affects the later choices
// however we can be greedy here because it is clear it is always better to replace
// with a small number whereas with longest common subsequence for example it is not
// clear apriori which is better
// rather than the second for loop we could optimize further by doing binary search on
// the sorted slice to get O(n logn) time complexity

// complexity
// time = O(n^2)
// space = O(n)
func lengthOfLIS(nums []int) int {
	sortedSubseq := []int{nums[0]}
	for _, num := range nums[1:] {
		// we can extend seq
		if num > sortedSubseq[len(sortedSubseq)-1] {
			sortedSubseq = append(sortedSubseq, num)
			// we can have same seq length but with lower initial number
		} else {
			i := 0
			for num > sortedSubseq[i] {
				i += 1
			}
			sortedSubseq[i] = num
		}
	}
	return len(sortedSubseq)
}
