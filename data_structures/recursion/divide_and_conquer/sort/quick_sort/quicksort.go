package divide_and_conquer

import (
	"math/rand"
	"sort"
)

// standard partition in place swap
// in order to use generic comparison requires a comparison function Less
// which implements the sort.Interface
func partition(s sort.Interface, start int, end int, pivotIndex int) int {
	s.Swap(start, pivotIndex) // move it to the beginning
	i := start + 1
	j := end
	for i <= j {
		for i <= end && s.Less(i, start) { // pivot is now at start
			i++
		}
		for j >= start && s.Less(start, j) { // pivot is now at start
			j--
		}
		if i <= j {
			s.Swap(i, j)
			i++
			j--
		}
	}
	s.Swap(start, j)
	return j
}

func sortRecursiveWithSwap(s sort.Interface, start int, end int) {
	if start >= end {
		return
	}
	pivotIndex := partition(s, start, end, rand.Intn(end-start+1)+start)
	// we use a recursive closure rather than passing slices
	// this is because although slices are cheap (they do not copy the underlying array)
	// there is some overhead in constructing the slice object
	sortRecursiveWithSwap(s, start, pivotIndex+1)
	sortRecursiveWithSwap(s, pivotIndex+1, end)
}

func sortRecursiveInPlace(s sort.Interface) {
	sortRecursiveWithSwap(s, 0, s.Len()-1)
}
