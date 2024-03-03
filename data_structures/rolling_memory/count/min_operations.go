package data_structures

// approaches
// * just loop through and count the number of numbers less than k
func minOperations(nums []int, k int) int {
	count := 0
	for _, num := range nums {
		if num < k {
			count++
		}
	}
	return count
}
