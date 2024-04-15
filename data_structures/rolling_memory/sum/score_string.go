package data_structures

func scoreOfString(s string) int {
	cumScore := 0
	for i := 1; i < len(s); i++ {
		cumScore += abs(int(s[i-1]), int(s[i]))
	}
	return cumScore
}

func abs(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}
