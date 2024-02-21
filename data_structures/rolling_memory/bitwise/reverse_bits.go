package data_structures

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res << 1 // shift bits left by 1
		res = res | (num & 1)
		num = num >> 1
	}
	return res
}
