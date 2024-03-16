package data_structures

import (
	"container/heap"
	"sort"
)

// clarifying questions
// * do we need to calculate lexicographic? probably not - can just choose smallest letter is possible
// * for the delta we don't actually care about positioning? if we have an a whether we choose a earlier or later still adds +cost
// * suppose have a, _, a, a, _ <- if we add a a new to first position that has cost(1) to it, and adds 2 costs to later so +3, if add to later position that it has +3 so where we put the new letter doesn't matter?
// * what is the cost function if single a it is 0, if two as it is 1, a, a, a = 2, 4as = 6? i.e. it is the sum from 0 to n - 1 which is (n - 1)(n) / 2  if n = 4 then 3 * 4 / 2 = 6

// approaches
// * map of letters and costs where choose the smallest cost everytime,
// * create a heap where minimization function is 1. smallest cost in string so far 2.
func minimizeStringValue(s string) string {
	// setup map
	m := make([]int, 26) // map of counts

	// go through string excluding ?s and add to heap
	questionMarkI := make([]int, 0)
	for i, c := range s {
		if c == '?' {
			questionMarkI = append(questionMarkI, i)
		} else {
			m[int(c-'a')]++
		}
	}
	// setup heap
	h := make(Heap, 0)
	heap.Init(&h)
	for i, count := range m {
		count = count + 1
		cost := cost(count) // add one as this is the marginal cost
		letter := rune('a' + i)
		heap.Push(&h, data{cost: cost, count: count, letter: letter})
	}
	// go through ?s and add to heap
	qMarks := make([]rune, len(questionMarkI))
	for i, _ := range qMarks {
		popped := heap.Pop(&h).(data)
		qMarks[i] = popped.letter
		popped.count++
		popped.cost = cost(popped.count)
		heap.Push(&h, popped)
	}
	sort.Slice(qMarks, func(i, j int) bool {
		return qMarks[i] < qMarks[j]
	})
	sRunes := []rune(s)
	for i, q := range questionMarkI {
		sRunes[q] = qMarks[i]
	}
	return string(sRunes)
}

type data struct {
	cost   int
	count  int
	letter rune
}

func cost(n int) int {
	if n == 0 {
		return 0
	}
	return n * (n - 1) / 2
}

type Heap []data

func (h *Heap) Push(x interface{}) {
	(*h) = append(*h, x.(data)) // cast
}
func (h *Heap) Pop() interface{} {
	previous, popped, n := *h, data{}, h.Len()
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Heap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].letter < h[j].letter
	}
	return h[i].cost < h[j].cost
}
