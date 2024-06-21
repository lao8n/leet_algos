package data_structures

// use kahn's topological sorting algorithm
// basically: queue + in-degree map
func alienOrder(words []string) string {
	// in-degree map - necessary so we have all characters
	inDegreeMap := make(map[rune]int, 0)
	for _, word := range words {
		for _, c := range word {
			// new character
			inDegreeMap[c] = 0 // no need for if statement - just override
		}
	}

	// adjacency list
	adjList := make(map[rune][]rune, 0) // out-degree mapping
	for i := 1; i < len(words); i++ {
		prevWord, currWord := words[i-1], words[i]
		// isolate the reason for the ordering for each pair
		minLength := len(prevWord)
		if len(currWord) < minLength {
			minLength = len(currWord)
		}
		// no useful information if difference is lengths
		foundReason := false
		for j := 0; j < minLength; j++ {
			prevRune, currRune := rune(prevWord[j]), rune(currWord[j])
			if prevRune != currRune {
				// prevWord[j] comes before currWord[j]
				adjList[prevRune] = append(adjList[prevRune], currRune)
				// currWord rune has + 1 in-degree
				inDegreeMap[currRune]++
				foundReason = true
				break // found reason for difference
			}
		}
		if !foundReason && len(prevWord) > len(currWord) {
			return ""
		}
	}

	// setup queue
	queue := make([]rune, 0)
	for c, d := range inDegreeMap {
		if d == 0 {
			queue = append(queue, c)
		}
	}

	solution := make([]rune, 0)
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		// process popped
		solution = append(solution, popped)
		// add neighbours to queue
		for _, neighbour := range adjList[popped] {
			inDegreeMap[neighbour]--
			if inDegreeMap[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	if len(solution) != len(inDegreeMap) { // missing characters
		return ""
	}
	return string(solution)
}
