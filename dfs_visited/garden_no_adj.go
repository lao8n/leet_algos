package data_structures

import "fmt"

// clarifying questions
// * n gardens - bidirectional path from x <-> y
// * at most 3 paths?
// * 4 flowers - different flowers

// approaches
// * seems like the colouring problem
// * only need a single approach so could 1. build adjaceny matrix 2. recurse through shape trying different colours 3. backtrack if doesn't lead to a solution
// * don't think it is a dp for example rob houses because although interconnected decision not clear what the state would be set of colours for some houses?
// * thought it might be the timetable problem where start with zero dependency but probably different...
func gardenNoAdj(n int, paths [][]int) []int {
	// create adj matrix
	adjMatrix := make(map[int][]int) // could use set but should be unique
	for _, path := range paths {
		x, y := path[0], path[1]
		adjMatrix[x] = append(adjMatrix[x], y)
		adjMatrix[y] = append(adjMatrix[y], x)
	}
	// recurse through
	flowers := make(map[int]int) // makes it easier to tell when all done
	solution := gardenNoAdjRec(1, flowers, n, adjMatrix)
	fmt.Println(flowers)
	return solution
}

// accumulate possible path
func gardenNoAdjRec(garden int, flowers map[int]int, n int, adjMatrix map[int][]int) []int {
	// base case
	if len(flowers) == n {
		solution := make([]int, n)
		for g, f := range flowers {
			solution[g-1] = f
		}
		return solution
	}

	if _, ok := flowers[garden]; ok {
		// been here before
		return []int{}
	}

	// recursive case
	// try to find all valid flowerings for current map
	usedColours := make(map[int]bool)
	for _, neighbour := range adjMatrix[garden] {
		// look up colours already there
		if flower, ok := flowers[neighbour]; ok {
			usedColours[flower] = true
		}
	}
	// fmt.Println(garden, flowers)
	for i := 1; i <= 4; i++ {
		if !usedColours[i] {
			// try this colour
			flowers[garden] = i
			solution := gardenNoAdjRec(garden+1, flowers, n, adjMatrix)
			if len(solution) > 0 {
				return solution
			}
			// backtrack
			// delete(flowers, garden)
		}
	}
	return []int{}
}
