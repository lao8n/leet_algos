package data_structures

import (
	"sort"
	"strings"
)

// clarifying questions
// * happy -> cheerful even though no direct mapping

// approaches
// * loop through text replacing with synonyms -> how have mapping? set?

// specifics
// * want to be able to go through text and for each word -> list of synonyms
// * text is only 10 words
func generateSentences(synonyms [][]string, text string) []string {
	// union find of synonyms
	uf := disjointSet{root: make(map[string]string)}
	// init
	for _, synonym := range synonyms {
		uf.root[synonym[0]] = synonym[0]
		uf.root[synonym[1]] = synonym[1]
	}
	// union
	for _, synonym := range synonyms {
		uf.union(synonym[0], synonym[1])
	}
	// flip map
	rootMap := make(map[string][]string)
	for word, _ := range uf.root {
		root := uf.find(word)
		rootMap[root] = append(rootMap[root], word)
	}
	// rootMap: map[cheerful:[cheerful happy joy] sorrow:[sorrow sad]]

	// replace each word with variations
	solutions := make([][]string, 1)
	solutions[0] = strings.Split(text, " ")
	for i, word := range strings.Split(text, " ") {
		newSolutions := [][]string{} // similar to letter combinations with digit letter map
		if root, ok := uf.root[word]; ok {
			for _, s := range solutions {
				for _, synonym := range rootMap[root] {
					s[i] = synonym
					newS := make([]string, len(s))
					copy(newS, s)
					newSolutions = append(newSolutions, newS)
				}
			}
		}
		if len(newSolutions) != 0 {
			solutions = newSolutions
		}
	}

	// convert slices of strings to strings
	solutionsStrings := make([]string, len(solutions))
	for i, s := range solutions {
		solutionsStrings[i] = strings.Join(s, " ")
	}

	// sort lexicographically
	sort.Strings(solutionsStrings)

	return solutionsStrings
}

type disjointSet struct {
	root map[string]string
	// don't worry about union by rank
}

func (uf *disjointSet) find(x string) string {
	if x == uf.root[x] {
		return x
	}
	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

// don't both with union by rank as have < 10
func (uf *disjointSet) union(x, y string) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		uf.root[x] = uf.root[y] // arbitrarily choose one
	}
}
