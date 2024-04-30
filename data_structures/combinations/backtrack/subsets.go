package data_structures

// clarifying questions
// * power set doesn't care about permutations correct?

// approaches
// * with combination questions we build a path, with permutation questions we build all possible paths from remaining nums
// * this is just a combinations question but we append very subpath - not just the final path to the solution

// implementation specifics
// * need to be careful with slices because they are pointers - usually we need to backtrack but copying every subpath so could avoid..
// formula for size of power set?
func subsets(nums []int) [][]int {
	d := data{
		nums:     nums,
		solution: make([][]int, 0),
	}
	d.backtrack([]int{}, 0)
	return d.solution
}

type data struct {
	nums     []int
	solution [][]int
}

func (d *data) backtrack(path []int, start int) {
	// base case
	if len(path) > len(d.nums) {
		return // without adding anything - we've recursed too far
	}
	copyPath := make([]int, len(path))
	copy(copyPath, path)
	d.solution = append(d.solution, copyPath)

	// recursive case
	for i := start; i < len(d.nums); i++ {
		path = append(path, d.nums[i])
		d.backtrack(path, i+1)
		path = path[:len(path)-1]
	}
}
