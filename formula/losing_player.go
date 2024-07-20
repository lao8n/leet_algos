package data_structures

// approach
// * every turn pick up 115 - only way to make 115 is 1 75 and 4 10s
func losingPlayer(x int, y int) string {
	turns75 := x
	turns10 := y / 4
	turns := min(turns75, turns10)
	if turns%2 == 0 {
		return "Bob"
	}
	return "Alice"
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
