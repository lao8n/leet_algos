package data_structures

// approaches
// * manually remove elements
// * count number of letters of each type - where can remove 2 at a time as long as >= 3 of them
func minimumLength(s string) int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	sum := 0
	for _, count := range counts {
		// if have 5 -> remove 2 then go to 3 -> got 1
		// if have 4 -> remove 2 then get stuck at 2
		// if have 3 -> remove 2 then get stuck at 1
		if count%2 == 0 {
			sum += 2 // has to be at least 2
		} else {
			sum += 1
		}
	}
	return sum
}
