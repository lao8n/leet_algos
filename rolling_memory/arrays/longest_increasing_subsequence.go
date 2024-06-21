package data_structures

// clarifying questions
// * has to be strictly increasing

// approaches
// * stack -> rolling max and length of subsequence...

// specifics
// * at any given point you have a choice of multiple subsequences each of length n and with max m.
func lengthOfLIS(nums []int) int {
	seq := make([]int, 1) // this is the best seq we can have with lowest number at each length
	seq[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		// can we make seq longer?
		if nums[i] > seq[len(seq)-1] {
			seq = append(seq, nums[i])
		} else {
			// can we make any seq of the same length but lower last value?
			j := 0
			for nums[i] > seq[j] { // can improve this section with binary search
				j++
			}
			seq[j] = nums[i] //
		}
	}
	return len(seq)
}
