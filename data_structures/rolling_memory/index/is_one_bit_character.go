package data_structures

// approaches
// * don't need dp because can greedily say if 0 bit then it increment, if a 1 increment by 2
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits)-1
}
