package data_structures

// approach
// * can't think of a way to do it with xor itself.. if just xor everything then will get product of single digits
func duplicateNumbersXOR(nums []int) int {
	counts := make(map[int]int)
	xor := 0
	for _, num := range nums {
		counts[num]++
		if counts[num] == 2 {
			xor ^= num
		}
	}
	return xor
}
