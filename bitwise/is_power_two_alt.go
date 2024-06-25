package data_structures

// key idea
// * power of two in binary representation is (0000001), (00001000), (0100000) i.e. only 1 bit is 1 all others are zero
// * otherwise there are multiple 1s
// * exception is 0 which is all 0s

// approaches
// * get right-most 1-bit: x & (-x) - works because of 2s complement where -x = not x + 1. for power of 2 x & (-x) == x but it will not be for non-power of 2
// * turn-off the right-most 1-bit: x & (x - 1)
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0 // turns off right-most one bit
}
