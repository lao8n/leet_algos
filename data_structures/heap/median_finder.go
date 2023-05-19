package data_structures

import (
	"container/heap"
	"math"
)

// maintain a heap but not of size n but of size n / 2
// won't work because what if suddenly get lots of small numbers - we need to track everything
// but then how handle popping all the elements and then pushing them back again every time we need the median?
// maybe we can exploit the underlying array but cannot just take n/2 array element
// instead can maintain two heaps a max and a min where we transfer elements between them

// choice between whether we rebalance the heaps after every add number of only when finding median
type MedianFinder struct {
	belowMedian MaxHeap
	aboveMedian MinHeap
}

func Constructor() MedianFinder {
	maxHeap := make(MaxHeap, 0)
	minHeap := make(MinHeap, 0)
	return MedianFinder{
		belowMedian: maxHeap,
		aboveMedian: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// main case
	belowMedianMax, aboveMedianMin := math.MinInt, math.MaxInt
	if this.belowMedian.Len() > 0 {
		belowMedianMax = this.belowMedian.Peek()
	}
	if this.aboveMedian.Len() > 0 {
		aboveMedianMin = this.aboveMedian.Peek()
	}
	if num < belowMedianMax {
		heap.Push(&this.belowMedian, num)
	} else if num > aboveMedianMin {
		heap.Push(&this.aboveMedian, num)
		// in between so push to smaller heap
	} else if this.belowMedian.Len() < this.aboveMedian.Len() {
		heap.Push(&this.belowMedian, num)
	} else {
		heap.Push(&this.aboveMedian, num)
	}

	// rebalance
	// belowMedian too large
	for this.belowMedian.Len()-this.aboveMedian.Len() > 1 {
		popped := heap.Pop(&this.belowMedian)
		heap.Push(&this.aboveMedian, popped)
	}
	for this.aboveMedian.Len()-this.belowMedian.Len() > 1 {
		popped := heap.Pop(&this.aboveMedian)
		heap.Push(&this.belowMedian, popped)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.aboveMedian.Len() == this.belowMedian.Len() {
		above, below := this.aboveMedian.Peek(), this.belowMedian.Peek()
		return float64(above+below) / 2
	} else if this.aboveMedian.Len() > this.belowMedian.Len() {
		return float64(this.aboveMedian.Peek())
	} else {
		return float64(this.belowMedian.Peek())
	}
}

type MaxHeap []int

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h MaxHeap) Peek() int          { return h[0] }
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type MinHeap []int

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h MinHeap) Peek() int          { return h[0] }
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
