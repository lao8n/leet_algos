package data_structures

import "sort"

// clarifying questions
// * arr2 - distinct all in arr2 in arr1
func relativeSortArray(arr1 []int, arr2 []int) []int {
	// get counts of arr1
	map1 := make(map[int]int)
	for _, num := range arr1 {
		map1[num]++
	}
	// loop arr2 with counts from arr1
	solution := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		for i := 0; i < map1[num]; i++ {
			solution = append(solution, num)
		}
		delete(map1, num)
	}
	// add any leftovers from arr1
	leftovers := []int{}
	for num, count := range map1 {
		for i := 0; i < count; i++ {
			leftovers = append(leftovers, num)
		}
	}
	sort.Ints(leftovers)

	return append(solution, leftovers...)
}
