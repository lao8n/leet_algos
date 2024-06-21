package data_structures

func findWords(board [][]byte, words []string) []string {
	// build trie
	root := &node{children: make(map[byte]*node)}
	for _, word := range words {
		n := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if next, ok := n.children[c]; ok {
				n = next
			} else {
				n.children[c] = &node{children: make(map[byte]*node)}
				n = n.children[c]
			}
		}
		n.word = word
	}

	// recurse through trie
	d := data{
		board:  board,
		result: map[string]struct{}{},
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			d.recurse(i, j, root)
		}
	}

	// get result
	result := make([]string, len(d.result))
	i := 0
	for word, _ := range d.result {
		result[i] = word
		i++
	}

	return result
}

type node struct {
	word     string
	children map[byte]*node
}

type data struct {
	board  [][]byte
	result map[string]struct{}
}

func (d *data) recurse(i, j int, parent *node) {
	c := d.board[i][j]
	n := parent.children[c]
	if n == nil {
		return
	}

	// found a word
	if n.word != "" {
		d.result[n.word] = struct{}{}
		n.word = ""
	}

	// mark visited
	d.board[i][j] = '#'

	// explore neighbours
	for _, neighbour := range d.neighbours(i, j) {
		d.recurse(neighbour[0], neighbour[1], n)
	}

	// backtrack
	d.board[i][j] = c

	// optimization could incrementally remove the matched leaf node in trie

}

func (d *data) neighbours(i, j int) [][]int {
	m, n := len(d.board), len(d.board[0])
	trans := []int{1, 0, -1, 0, 1}
	neighbours := make([][]int, 0)
	for t := 1; t < len(trans); t++ {
		dx, dy := trans[t-1], trans[t]
		nx, ny := i+dx, j+dy
		if nx >= 0 && nx < m && ny >= 0 && ny < n && d.board[nx][ny] != '#' {
			neighbours = append(neighbours, []int{nx, ny})
		}
	}
	return neighbours
}
