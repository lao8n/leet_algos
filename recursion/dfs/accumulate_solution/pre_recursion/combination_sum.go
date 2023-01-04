package recursion

// two sum etc. = 1. two pointer approach 2. map to hit diff
// this is different = 1. want all combinations (not just one) 2. can repeat numbers

// choices
// 1. how collect solutions? pointer or return? return
// 2. how to acc combination? accumulator? acc arg
// 3. validate yourself or next? next
// 4. how to get unique? golang map keys cannot be slice and maps are how sets are defined
// 5. how to handle pointer to combination? make a copy if we want to keep the solution
func combinationSum(candidates []int, target int) [][]int {
	return combinationSumRecursive(candidates, target, []int{})
}

func combinationSumRecursive(candidates []int, remainingTarget int, combination []int) [][]int {
	// base case
	if remainingTarget == 0 {
		// need to copy here because otherwise updating pointer to same slice
		cpCombination := make([]int, len(combination))
		copy(cpCombination, combination)
		return [][]int{cpCombination}
	}
	if remainingTarget < 0 {
		x := [][]int{}
		return x
	}
	solutions := [][]int{}
	for i, c := range candidates {
		cSolutions := combinationSumRecursive(candidates[i:], remainingTarget-c, append(combination, c))
		solutions = append(solutions, cSolutions...)
	}
	return solutions
}
