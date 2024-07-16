package data_structures

// approaches
// * brute force - try every empty seat and count distance left and right to those seats
// * loop over and count empty seats in a row and / 2 - but also need edges
func maxDistToClosest(seats []int) int {
	maxEmpty := 0
	leftEmpty, leftEmptyFlag := 0, true
	emptyCount := 0
	for _, seat := range seats {
		if seat == 1 {
			emptyCount = 0
			leftEmptyFlag = false
		} else {
			emptyCount++
			if leftEmptyFlag {
				leftEmpty++
			}
		}
		if emptyCount > maxEmpty {
			maxEmpty = emptyCount
		}
	}
	dist := max(leftEmpty, emptyCount)
	dist = max(dist, (maxEmpty+1)/2)
	return dist
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// (3 + 1) / 2 = 2
// (2 + 1) / 2 = 1
