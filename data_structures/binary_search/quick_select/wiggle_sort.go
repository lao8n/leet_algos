package data_structures

// https://en.wikipedia.org/wiki/Median_of_medians
// approaches
// * if we can partition nums into two camps: bigger than median and less than median then it automatically work - but to find median costs O(n logn) time if use heap, could use median of median
// * or could just have increasing and swap every pair
// * single pass swap numbers
func wiggleSort(nums []int) {
	median := (len(nums)+1)/2 - 1
	medianIndex := quickSelect(nums, 0, len(nums)-1, median)
	l := len(nums) - 1
	if l%2 == 1 { // this is odd index
		l--
	}
	wiggle := make([]int, len(nums))
	// smaller numbers - largest to smallest
	for i := 0; i <= medianIndex; i++ {
		s := medianIndex - i
		wiggle[i*2] = nums[s]
	}
	for i := 0; i <= medianIndex; i++ {
		s := len(nums) - 1 - i
		if i*2+1 < len(wiggle) {
			wiggle[i*2+1] = nums[s]
		}
	}
	copy(nums, wiggle)
}

func quickSelect(nums []int, l, r, mid int) int {
	for {
		if l == r {
			return l
		}
		partitionIndex := medianOfMedians(nums, l, r) // clever selection of pivot index
		pivotIndex := partition(nums, l, r, mid, partitionIndex)
		if mid == pivotIndex {
			return mid
		} else if mid < pivotIndex {
			r = pivotIndex - 1
		} else {
			l = pivotIndex + 1
		}
	}
	return 0
}

func partition(nums []int, l, r, mid, partitionIndex int) int {
	pivotValue := nums[partitionIndex]
	// move pivot Value to end
	nums[partitionIndex], nums[r] = nums[r], nums[partitionIndex]

	// move all elements less than pivot to the left of the pivot
	lI := l
	for i := l; i < r; i++ {
		if nums[i] < pivotValue {
			nums[i], nums[lI] = nums[lI], nums[i]
			lI++
		}
	}
	// move all elements equal to the pivot incl. the pivot
	eI := lI
	for i := lI; i <= r; i++ {
		if nums[i] == pivotValue {
			nums[i], nums[eI] = nums[eI], nums[i]
			eI++
		}
	}
	// where is our desired number?
	if mid < lI {
		return lI
	} else if mid <= eI {
		return mid
	} else {
		return eI
	}
}

func medianOfMedians(nums []int, l, r int) int {
	// <= five elements
	if r-l < 5 {
		return medianOf5(nums, l, r)
	}
	// move medians of five - element subgroup to the first n / 5 positions
	for i := l; i < r; i += 5 {
		subRight := i + 4
		if subRight > r {
			subRight = r
		}
		median5 := medianOf5(nums, l, subRight)
		ithGroup := l + (i-l)/5
		nums[median5], nums[ithGroup] = nums[ithGroup], nums[median5]
	}

	// mutual recursion
	midMedians := l + 1 + (r-l)/10 // why plus 1?
	rightMedians := l + (r-l)/5
	return quickSelect(nums, l, rightMedians, midMedians)
}

func medianOf5(nums []int, l, r int) int {
	// insertion sort - sorted and unsorted
	for i := l + 1; i < r; i++ { // unsorted section
		j := i
		for j > l && nums[j-1] > nums[j] { // swap into sorted section
			nums[j-1], nums[j] = nums[j], nums[j-1]
			j--
		}
	}
	return l + (r-l)/2
}
