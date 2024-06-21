package data_structures

import "container/heap"

// Heap interface
// type Interface interface {
// 	sort.Interface
// 	Push(x any) // add x as element Len()
// 	Pop() any   // remove and return element Len() - 1.
// }

type wordFreq struct {
	word string
	freq int
}

type wordFreqs []wordFreq

func (wf *wordFreqs) Push(x interface{}) {
	*wf = append(*wf, x.(wordFreq))
}

func (wf *wordFreqs) Pop() interface{} {
	listWordFreqs := *wf
	n := len(listWordFreqs)
	*wf = listWordFreqs[:n-1]
	return listWordFreqs[n-1]
}

func (wf wordFreqs) Swap(i, j int) {
	wf[i], wf[j] = wf[j], wf[i]
}

func (wf wordFreqs) Len() int { return len(wf) }

func (wf wordFreqs) Less(i, j int) bool {
	if wf[i].freq == wf[j].freq {
		return wf[i].word < wf[j].word
	}
	return wf[i].freq > wf[j].freq
}

func topKFrequent(words []string, k int) []string {
	wordFreqSlice := make(wordFreqs, 0, k+1)
	wordFreqMap := make(map[string]int, len(words))

	// create map of freqs O(n)
	for _, word := range words {
		wordFreqMap[word] += 1
	}

	// push words onto heap O(n) * O(logn)
	for key, value := range wordFreqMap {
		heap.Push(&wordFreqSlice, wordFreq{key, value})
		// max heap length
		if len(wordFreqSlice) > k && wordFreqSlice[0].freq < value {
			heap.Pop(&wordFreqSlice)
		}
	}

	// make output list O(n)
	output := make([]string, k)
	for i := 0; i < k; i++ {
		poppedWord := heap.Pop(&wordFreqSlice).(wordFreq)
		output[i] = poppedWord.word
	}
	return output
}
