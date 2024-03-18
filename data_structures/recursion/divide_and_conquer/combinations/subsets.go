package data_structures

// choice:
// 1. recurse backwards from [] base case, looping over solutions adding subsets to return
// 2. recurse forwards with accumulator building to subset, looping over solutions adding subsets to return statement or as pointer
// approach = subsets is basically the combinations approach but appending solutions rather than adding
// avoid duplicates = only loop over the solutions or the subset - easiest is one value at a time

func subsets(nums []int) [][]int {
	// base case
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	// recursive case - loop over solutions but only add one value at a time
	solutions := subsets(nums[:len(nums)-1])
	for _, s := range solutions {
		newS := make([]int, len(s))
		copy(newS, s)
		solutions = append(solutions, append(newS, nums[len(nums)-1]))
	}
	return solutions
}
