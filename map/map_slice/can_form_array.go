package data_structures

// approaches
// * use set to match - but doesn't work with pieces..
// * use a set of pieces - then greedily match pieces - key could be first num which is unique
// * could sort but doesn't help necessarily

// specifics
// * what is have array elements at the end
// * what is have unused pieces at the end?
func canFormArray(arr []int, pieces [][]int) bool {
	mapPieces := make(map[int][]int)
	for _, piece := range pieces {
		mapPieces[piece[0]] = piece // 16 -> [16, 18, 49]
	}
	i := 0
	for i < len(arr) {
		num := arr[i]
		if piece, ok := mapPieces[num]; ok {
			i++                               // we start by comparing next along as already compared
			for j := 1; j < len(piece); j++ { // already matched first element
				if i >= len(arr) {
					return false
				}
				if arr[i] != piece[j] {
					return false
				}
				i++
			}
		} else {
			return false // no piece found
		}
	}
	return true
}
