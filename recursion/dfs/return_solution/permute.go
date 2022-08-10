package recursion

// how to deal with repetitions? shouldn't be a problem as ordering does matter
// how to track which numbers used? maybe arg of indices left
// how to acc solutions? either pointer or in return -> try return
// need acc function? probs not
// how to add to solutions? first item cannot append if nothing htere
func permute(nums []int) [][]int {
	// base case
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	// recursive case
	solutions := [][]int{}
	for i, n := range nums {
		// get all nums except n
		lessThanINums := []int{}
		if i >= 0 {
			lessThanINums = make([]int, len(nums[:i]))
			copy(lessThanINums, nums[:i])
		}
		moreThanINums := []int{}
		if i+1 < len(nums) {
			moreThanINums = make([]int, len(nums[i+1:]))
			copy(moreThanINums, nums[i+1:])
		}
		withoutN := append(lessThanINums, moreThanINums...)

		// find all their permutations
		solutionsWithoutN := permute(withoutN)

		// combine them with n
		solutionsWithN := make([][]int, len(solutionsWithoutN))
		for j, w := range solutionsWithoutN {
			solutionsWithN[j] = append(w, n)
		}
		solutions = append(solutions, solutionsWithN...)
	}
	return solutions
}
