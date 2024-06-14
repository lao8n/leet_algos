package data_structures

// approaches
// * number of pairs where they are equal - map of counts? -> then convert into
func numEquivDominoPairs(dominoes [][]int) int {
	pairs := make(map[[2]int]int, 0)
	count := 0
	for _, domino := range dominoes {
		d := sortDomino(domino)
		if c, ok := pairs[d]; ok {
			count += c
		}
		pairs[d]++
	}
	return count
}

func sortDomino(domino []int) [2]int {
	min, max := domino[0], domino[1]
	if min > max {
		min, max = max, min
	}
	return [2]int{min, max} // unique ordering
}
