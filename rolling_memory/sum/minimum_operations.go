package data_structures

// clarifying questions
// * add or subtract means guaranteed to make divisible by 3
// approaches
// * just need to count num of non 3 and then one operation each
func minimumOperations(nums []int) int {
	count := 0
	for _, num := range nums {
		if num%3 != 0 {
			count++
		}
	}
	return count
}
