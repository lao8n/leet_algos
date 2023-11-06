package data_structures

// clarifying questions
// * already sorted - k closest integers to x
// * final output? should be sorted... k can be very large - length of array
// approaches
// * two pointers -> loop through arr to find point closest to x and for loop from that point
// * can optimise initial search with binary search

// specifics
// * how handle two elements either side of x vs exact match of x?
// * how output answers - want to avoid needing to sort - maybe two slices append and then reverse one onto the other?
func findClosestElements(arr []int, k int, x int) []int {
	l, r := findX(arr, x)
	i, j := 0, 0
	// add elements to dec & inc
	for i+j < k {
		// pick lower number
		if l >= 0 && r < len(arr) {
			if closerTo(x, arr[l], arr[r]) {
				i++
				l--
			} else {
				j++
				r++
			}
			// or pick any valid number
		} else if l >= 0 {
			i++
			l--
		} else if r < len(arr) {
			j++
			r++
		} else {
			break // no valid number
		}
	}
	return arr[l+1 : l+1+k]
}

func findX(arr []int, x int) (int, int) {
	l, r := 0, len(arr)-1
	for r-l > 1 {
		// find midpoint
		m := (l + r) / 2
		// pick side
		if arr[m] < x { // go right
			l = m
		} else { // go left
			r = m
		}
	}
	return l, r
}

// a closer to x than b
func closerTo(x int, a int, b int) bool {
	ax := a - x
	if ax < 0 {
		ax = -ax
	}
	bx := b - x
	if bx < 0 {
		bx = -bx
	}
	return ax <= bx
}
