package data_structures

// approaches
// * dp as one decision affects later decisions
// * greedy approach always choose smallest number that is inside?
// * dp top down -> with memo? or with path?

// specifics
// * cannot have slice as key in map, so convert to a string
// * need to take copies of slices or backtrack
// * problem is we have too many states - need to reduce the number of states so rather than representing by array l & r array instead
// * replace ithTotal := d.recursive(append(leftPart, rightPart...))
func maxCoins(nums []int) int {
	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(nums))
	}
	d := data{
		nums: nums,
		memo: memo,
	}
	return d.recursive(0, len(nums)-1)
}

type data struct {
	nums []int
	memo [][]int
}

func (d *data) recursive(l, r int) int {
	// base case
	if l > r {
		return 0
	}
	// memoized
	if d.memo[l][r] != 0 {
		return d.memo[l][r]
	}

	// recursive
	// partition on the last to burst, as this makes partition valid
	bestTotal := 0
	for i := l; i <= r; i++ {
		// handle edges
		lessNeighbour := 1
		if l > 0 {
			lessNeighbour = d.nums[l-1]
		}
		moreNeighbour := 1
		if r < len(d.nums)-1 {
			moreNeighbour = d.nums[r+1]
		}
		// valid calculation if d.nums is the last to burst
		newTotal := d.nums[i] * lessNeighbour * moreNeighbour
		// divide and conquer
		ithTotal := d.recursive(l, i-1) + d.recursive(i+1, r)
		if ithTotal+newTotal > bestTotal {
			bestTotal = ithTotal + newTotal
		}
	}
	d.memo[l][r] = bestTotal
	return bestTotal
}
