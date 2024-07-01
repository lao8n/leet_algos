// clarifying question
// * two neighbours both occupied or both vacant then cell becomes occupied
// * otherwise vacant.
// * handle cycle? yes potentially

// approach
// * manually update prison cells - but look for cycles
// * kind of a dp with memoization question
// specifics
// * how have key? could create a string - or could uniquely define wiht number
func prisonAfterNDays(cells []int, n int) []int {
	// cycle detection with map
	cycle := make(map[int]int) // key -> iteration
	// for loop apply rules
	i := 0
	for i < n {
		// process cur
		k := prisonCellKey(cells)
		if j, ok := cycle[k]; ok {
			// we have seen these set before at index j
			cycleLength := i - j
			remainingCycles := (n - i) / cycleLength // 80 / 5
			i = i + remainingCycles*cycleLength
			if i == n {
				break
			}
		}
		// go to next
		updateValues(cells)
		cycle[k] = i
		i++
	}
	return cells
}

// doesn't matter which way we do it
func prisonCellKey(cells []int) int {
	power2 := 1
	key := 0
	for i := 0; i < len(cells); i++ {
		if cells[i] == 1 {
			key += power2
		}
		power2 *= 2
	}
	return key
}

// in place
func updateValues(cells []int) {
	prev := 0
	for i := 0; i < len(cells); i++ {
		if i == 0 || i == len(cells)-1 {
			prev = cells[i]
			cells[i] = 0
		} else {
			cur := cells[i]
			if prev == cells[i+1] {
				cells[i] = 1
			} else {
				cells[i] = 0
			}
			prev = cur
		}
	}
}