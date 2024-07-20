package data_structures

// clarifying questions
// * disjoint - not even start and finish at same point?

// approach
// * two pointers - f and s ->
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intersection := make([][]int, 0)
	fI, sI := 0, 0 // interval
	// both lists
	for fI < len(firstList) && sI < len(secondList) {
		fCur := firstList[fI]
		sCur := secondList[sI]
		openI := max(fCur[0], sCur[0])
		closeI := min(fCur[1], sCur[1])
		// fmt.Println(fI, sI, openI, closeI)
		if closeI >= openI { // valid interval
			intersection = append(intersection, []int{openI, closeI})

		}
		if fCur[1] < sCur[1] {
			fI++
		} else {
			sI++
		}
	}
	// ignore remaining because cannot be intersection
	return intersection
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
