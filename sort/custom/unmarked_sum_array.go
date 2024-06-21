package data_structures

import "sort"

// clarifying questions
// * n nums (positive integers)
// * m queries [index, k]
// * mark element at index, then mark small elements in nums
// * return unmarked elements

// approaches
// * want to sort elements and remove but also need to be able to remove at an index.. if use heap for sorting - doesn't help with index..
// * how get 3 smallest that are unmarked?
func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	sorted := make([]data, len(nums))
	used := make(map[int]bool)
	sum := int64(0)
	for i, num := range nums {
		sorted[i] = data{num: num, index: i}
		used[i] = false // index
		sum += int64(num)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].num == sorted[j].num {
			return sorted[i].index < sorted[j].index
		}
		return sorted[i].num < sorted[j].num
	})

	solution := []int64{}
	s := 0
	for _, query := range queries {
		index, k := query[0], query[1]
		// process index
		if !used[index] { // we haven't gone here yet
			sum -= int64(nums[index])
			used[index] = true
		} // else ignore
		// process mins
		j := 0
		for j < k && s < len(sorted) {
			// new number
			if !used[sorted[s].index] {
				sum -= int64(sorted[s].num)
				used[sorted[s].index] = true
				j++
			} // else ignore
			s++
		}
		// add to solution
		solution = append(solution, sum)
	}
	return solution
}

type data struct {
	num   int
	index int
}
