package data_structures

// clarifying questions
// * find index of occurrence of x in nums
// approaches
// * create a map of occurrences for each one you find
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	occurrences := make(map[int]int)
	o := 1
	for i, num := range nums {
		if num == x {
			occurrences[o] = i
			o++
		}
	}
	result := make([]int, len(queries))
	for i, query := range queries {
		if oi, ok := occurrences[query]; ok {
			result[i] = oi
		} else {
			result[i] = -1
		}
	}
	return result
}
