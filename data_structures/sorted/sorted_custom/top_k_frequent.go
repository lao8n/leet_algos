package data_structures

// approach 1
// map O(n) store frequencies
// sort keys by values O(k logk)
// sort keys by lexicographical value O(k logk)

// approach 2
// maybe use min-tree

// approach 3
// could also sort first then store in map

// consideration 1
// could we only store k most frequent in map
import "sort"

func topKFrequent(words []string, k int) []string {
	// create map O(n)
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	// get keys O(n)
	wordKeys := make([]string, len(freqMap))
	i := 0
	for wordKey, _ := range freqMap {
		wordKeys[i] = wordKey
		i++
	}

	// sort by values (O(n logn))
	sort.SliceStable(wordKeys, func(i, j int) bool {
		if freqMap[wordKeys[i]] > freqMap[wordKeys[j]] {
			return true
		} else if freqMap[wordKeys[i]] == freqMap[wordKeys[j]] {
			// frequencies are the same - so sort lexicographically
			if wordKeys[i] < wordKeys[j] {
				return true
			}
		}
		return false
	})

	return wordKeys[:k]
}
