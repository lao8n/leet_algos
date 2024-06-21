package data_structures

import "sort"

// clarifying questions
// * if same rank do you need to skip? no because 2 3s then 4 not 5

// approaches
// * brute force is O(n^n) where search for largest, then next largest etc.
// * start from top? problem is don't know what number to place (it isn't just the length of the array), better to start from min?
// * could sort list of index and number and then you know where each number should be... no point using heap because it is n logn if all digits.. stack doesn't help as isolating increasing or decreasing isn't useful
// * don't want map because although could have index cannot sort keys.. although could just look up indices (including multiple)
func arrayRankTransform(arr []int) []int {
	// create map of num -> [] indices
	// O(n)
	m := make(map[int][]int)
	for i, num := range arr {
		m[num] = append(m[num], i)
	}

	// create keys from map -> and sort
	// O(n)
	keys := make([]int, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	// O(n logn)
	sort.Ints(keys)

	// solution
	rank := 1
	solution := make([]int, len(arr))
	for _, k := range keys {
		for _, v := range m[k] {
			solution[v] = rank
		}
		rank++
	}

	return solution
}
