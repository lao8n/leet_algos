package data_structures

// approaches
// * validity usually test with a stack
// * bfs counter approach?
// * rather than having to test validity of each part instead have left & right counts for each index?
func removeInvalidParentheses(s string) []string {
	var res []string
	visited := make(map[string]bool)
	queue := []string{s}
	foundValid := false

	// BFS traversal
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if the current string is valid
		if isValid(current) {
			res = append(res, current)
			foundValid = true
		}

		if foundValid {
			continue // If a valid expression is found, no need to process further
		}

		// Generate all possible variations of the current string by removing one parenthesis at a time
		for i := 0; i < len(current); i++ {
			if i != 0 && current[i] == current[i-1] {
				continue // Skip duplicate characters to avoid duplicates in the result
			}
			if current[i] != '(' && current[i] != ')' {
				continue // Skip non-parenthesis characters
			}

			newStr := current[:i] + current[i+1:]

			// Check if the new string has been visited before
			if !visited[newStr] {
				visited[newStr] = true
				queue = append(queue, newStr)
			}
		}
	}

	return res
}

func isValid(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
