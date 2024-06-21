package data_structures

// approach
// * use bfs
// specifics
// * create intermediate state with *
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wl := len(beginWord)
	// adj list
	adjList := make(map[string][]string) // star word to word
	for _, word := range wordList {
		for i := 0; i < wl; i++ {
			wordByte := []byte(word)
			wordByte[i] = '*'
			starWord := string(wordByte)
			adjList[starWord] = append(adjList[starWord], word)
		}
	}
	// bfs with queue
	queue := make([]data, 1)
	queue[0] = data{word: beginWord, count: 1}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	for len(queue) > 0 {
		// update and check
		popped := queue[0]
		queue = queue[1:]
		visited[popped.word] = struct{}{}
		if popped.word == endWord {
			return popped.count
		}
		// add adjacent words
		for i := 0; i < wl; i++ {
			wordByte := []byte(popped.word)
			wordByte[i] = '*'
			starWord := string(wordByte)
			// loop over neighbours
			for _, neighbour := range adjList[starWord] {
				if _, ok := visited[neighbour]; !ok {
					queue = append(queue, data{word: neighbour, count: popped.count + 1})
				}
			}
		}
	}
	return 0
}

type data struct {
	word  string
	count int
}
