// approach 1 = 1. sort O(n logn) 2. merge O(n)
// approach 2 = 1. merge sort combined -> divide and conquer - basically we adapt
// the standard merge sort to also merge intervals at the same time rather than
// sorting and then
// question: is there a cleaner way than the ugly two pointer approach where you
// need to do <l & <r then clean up with <l & <r afterwards?
func merge(intervals [][]int) [][]int {
	// base case
	intervalsLength := len(intervals)
	if intervalsLength <= 1 {
		return intervals
	}
	// recursive case
	// divide
	leftIntervals := merge(intervals[:intervalsLength/2])
	rightIntervals := merge(intervals[intervalsLength/2:])
	// conquer
	leftPointer, rightPointer := 0, 0
	mergedIntervals := [][]int{}
	for leftPointer < len(leftIntervals) && rightPointer < len(rightIntervals) {
		var intervalToMerge []int
		if leftIntervals[leftPointer][0] < rightIntervals[rightPointer][0] {
			intervalToMerge = leftIntervals[leftPointer]
			leftPointer++
		} else {
			intervalToMerge = rightIntervals[rightPointer]
			rightPointer++
		}
		if len(mergedIntervals) == 0 {
			mergedIntervals = append(mergedIntervals, intervalToMerge)
		}
		// overlap
		if mergedIntervals[len(mergedIntervals)-1][1] >= intervalToMerge[0] {
			if mergedIntervals[len(mergedIntervals)-1][1] < intervalToMerge[1] {
				mergedIntervals[len(mergedIntervals)-1][1] = intervalToMerge[1]
			}
		} else {
			mergedIntervals = append(mergedIntervals, intervalToMerge)
		}
	}
	for leftPointer < len(leftIntervals) {
		// overlap
		if mergedIntervals[len(mergedIntervals)-1][1] >= leftIntervals[leftPointer][0] {
			if mergedIntervals[len(mergedIntervals)-1][1] < leftIntervals[leftPointer][1] {
				mergedIntervals[len(mergedIntervals)-1][1] = leftIntervals[leftPointer][1]
			}
			// no overlap
		} else {
			mergedIntervals = append(mergedIntervals, leftIntervals[leftPointer])
		}
		leftPointer++
	}
	for rightPointer < len(rightIntervals) {
		// overlap
		if mergedIntervals[len(mergedIntervals)-1][1] >= rightIntervals[rightPointer][0] {
			if mergedIntervals[len(mergedIntervals)-1][1] < rightIntervals[rightPointer][1] {
				mergedIntervals[len(mergedIntervals)-1][1] = rightIntervals[rightPointer][1]
			}
			// no overlap
		} else {
			mergedIntervals = append(mergedIntervals, rightIntervals[rightPointer])
		}
		rightPointer++
	}
	return mergedIntervals
}