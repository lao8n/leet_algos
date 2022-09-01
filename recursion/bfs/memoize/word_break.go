package recursion

// bfs
func wordBreak(s string, wordDict []string) bool {
	words := map[string]bool{}
	for _, word := range wordDict {
		words[word] = true
	}
	return bfs(s, words)
}

func bfs(s string, words map[string]bool) bool {
	var queue []int
	memoized := make([]bool, len(s))

	queue = append(queue, 0)
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if !memoized[start] {
			for end := start + 1; end <= len(s); end++ {
				if words[s[start:end]] {
					queue = append(queue, end)
					if end == len(s) {
						return true
					}
				}
			}
			memoized[start] = true
		}
	}
	return false
}
