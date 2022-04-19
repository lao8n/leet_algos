package data_structures

func sortedSquares(nums []int) []int {
	// square than sort = O(n) to square each element then O(n logn) to sort
	// two-pointers = O(n) to find smallest negative element and smallest positive element
	// then use two-pointers to go back along comparing as you build -> similar to merge
	// sort which is another O(n) for O(2n). the trick is the output order is not random,
	// it is determined by the inputs which are ordered themselves and we can exploit.
	// i.e. square -> sort = sort -> square

	// key idea = rather than start with the smallest and work out, start with the largest and work in
	n := len(nums)
	var l, r = 0, n - 1
	output := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		lSquared := nums[l] * nums[l]
		rSquared := nums[r] * nums[r]
		if lSquared > rSquared {
			output[i] = lSquared
			l++
		} else {
			output[i] = rSquared
			r--
		}
	}
	return output
}
