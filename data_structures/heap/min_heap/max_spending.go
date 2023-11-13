package data_structures

import "container/heap"


// clarifying questions
// * m shops - n items with more expensive items first in the row
// * m x n matrix so where does day come into it?
// approaches
// * values[i][j] * d - buy dth item if haven't already bought..
// * greedy approach buy the cheapest item first.. it is just day * largest value
// * just sort them all into one big array -> create a heap of values for each and pop
func maxSpending(values [][]int) int64 {
    h := make(Heap, 0)
    heap.Init(&h)
    
    // push initial values
    m, n := len(values), len(values[0])
    for i := 0; i < m; i++ {
        heap.Push(&h, data{row: i, col: n-1, val: values[i][n - 1]})
    }
    var rollingSum int64
    day := 1
    for h.Len() > 0 {
        popped := heap.Pop(&h).(data)
        rollingSum += int64(day * popped.val)
        if popped.col - 1 >= 0 {
            heap.Push(&h, data{row: popped.row, col: popped.col - 1, val: values[popped.row][popped.col - 1]})
        }
        day++
    }
    return rollingSum
}

type data struct {
    row int
    col int
    val int
}

type Heap []data
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(data))}
func (h *Heap) Pop() interface{} {
    previous, n, popped := *h, h.Len(), data{}
    *h, popped = previous[:n-1], previous[n-1]
    return popped
}
func (h Heap) Len() int { return len(h)}
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h Heap) Less(i, j int) bool {return h[i].val < h[j].val }