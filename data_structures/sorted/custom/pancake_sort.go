package data_structures

// clarifying questions
// * pancake 0 -> k but not after k?
// * largest number to the right

// approaches
// * inductive approach = k and above assume sorted. then want to flip largest element into 0th spot and then flip out to top.. i.e. two flips - similar to bubble sort where extract largest number and put it in the sorted section - however a flip can be done in O(k / 2) time so it is O(k) and for n numbers that is O(nk)
// * can we do better? so far we are focusing on just moving one element in at a time? we can't do a quick sort recursion type approach because only have 1-k flip
func pancakeSort(arr []int) []int {
	s := len(arr) - 1
	flips := make([]int, 0)
	for s > 0 {
		// find largest
		l := largest(arr, s)
		if l == s {
			s--
			continue
		}
		// swap to first index
		flip(arr, l)
		flips = append(flips, l+1) // k
		// swap to sth index
		flip(arr, s)
		flips = append(flips, s+1) // k
		s--
	}
	return flips
}

// find largest num
func largest(arr []int, k int) int {
	maxIndex := 0
	for i := 0; i <= k; i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// in place
func flip(arr []int, k int) {
	for i := 0; i <= k/2; i++ {
		arr[i], arr[k-i] = arr[k-i], arr[i]
	}
}
