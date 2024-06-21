package data_structures

import "sort"

// approaches
// * greedily try any queue - if not valid push to back of queue
// * custom sort - but unclear how to manage number above and below a number?
// * topological sort where 0 before but once added say [4, 0], [5, 0] we'd like to be able to then process [2, 2] as [2,0] etc. i think we can have a map of all numbers and when add to queue then decrement. need adjacency map however of all values below a certain value -> maybe just loop over remaining?
// * greedy - trick is to place the shortest person in the ith index
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		// sort by heights then index
		ithPerson, jthPerson := people[i], people[j]
		if ithPerson[0] == jthPerson[0] {
			return ithPerson[1] < jthPerson[1] // ascending k values
		}
		return ithPerson[0] > jthPerson[0] // descending height
	})
	result := make([][]int, 0, len(people))
	// [4, 4], [5, 0] [5, 2], [6, 1], [7, 0], [7, 1]
	for _, person := range people {
		inFront := person[1]
		result = append(result, person)
		copy(result[inFront+1:], result[inFront:])
		result[inFront] = person
	}
	return result
}
