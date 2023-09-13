package data_structures

// two approaches
// 1. go through intermediary of an actual number
// 2. use the array as is
func plusOne(digits []int) []int {
	carry, i := 1, len(digits)-1
	for i >= 0 && carry >= 0 {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
		i--
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
