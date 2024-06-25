// simple solution use a map with counts
// maybe we having a rolling sum? or divide by two somehow?
// maybe boyer moore maintain candidates?
// can't rely upon sum trick because not like it is a continuous numbers
// constant space cannot be a set
// sum can determine whether odd number or even so could do second pass where only check if odd or even but still O(n/2) space complexity
// maybe two pointers where take summation remove elements from both side
// sort is easy but O(n logn) - wherever the point is the sum of all elements on both sides is constant.
// [left] [num] [right] if num is unique then both sides will be
// bit manipulation
// xor = only true if they are arguments are different
// bitwise xor = a XOR 0 = a and a XOR a = 0 thus multiply everything together
func singleNumber(nums []int) int {
	bitwiseProduct := 0
	for _, num := range nums {
		bitwiseProduct ^= num
	}
	return bitwiseProduct
}