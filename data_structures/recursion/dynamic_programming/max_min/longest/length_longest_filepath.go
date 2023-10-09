package data_structures

import "strings"

// forward pass = accumulate path len, index & memoization
// backward pass = return max path len
// recursive top-down dynamic programming
func lengthLongestPath(input string) int {
	splittedStrings := strings.Split(input, "\n") // ["dir", "\tsubdir1", "\t\tfile1.ext"]
	memo := make(map[int]int)                     // folder depth -> path length
	return dfs(splittedStrings, 0, 0, memo)
}

func dfs(paths []string, pathIndex int, pathLen int, memo map[int]int) int {
	// base cases
	if pathIndex >= len(paths) {
		return pathLen
	}
	// process
	// remove and count ts
	cur := 0
	count := 0
	for _, c := range paths[pathIndex] {
		if c != '\t' {
			break
		}
		cur++
		count++
	}

	// recursive case
	// found a file
	foundFilePathLen := 0
	if strings.Contains(paths[pathIndex], ".") {
		foundFilePathLen = memo[cur-1] + len(paths[pathIndex][cur:])
	}
	// path len ends with '/'
	memo[cur] = memo[cur-1] + len(paths[pathIndex][cur:]) + 1
	if foundFilePathLen > pathLen {
		return dfs(paths, pathIndex+1, foundFilePathLen, memo)
	} else {
		return dfs(paths, pathIndex+1, pathLen, memo)
	}
}
