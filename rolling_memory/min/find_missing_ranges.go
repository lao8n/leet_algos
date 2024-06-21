package data_structures

// clarifying questions
// * nums is sorted and unique? yes
// * nums all within inclusive range
// approaches
// * loop over nums and create range candidates
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	ranges := make([][]int, 0)
	start := lower
	for _, num := range nums {
		// generic case [...3, 50] - also handles first case
		finish := num - 1
		if finish >= start {
			ranges = append(ranges, []int{start, finish})
		}
		start = num + 1
	}
	if upper >= start {
		ranges = append(ranges, []int{start, upper})
	}
	return ranges
}
