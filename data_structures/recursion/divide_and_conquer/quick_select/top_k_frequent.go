package data_structures

import (
	"fmt"
	"math/rand"
)

// quick select - average O(n)
// example of conquer than divide (as opposed to divide and conquer)
// key points left, right, pivot index
func topKFrequent(nums []int, k int) []int {
	// create map of frequencies O(n)
	freqMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]] += 1 // unprotected map access ok
	}

	// create slice O(n)
	freqSlice := make([]data, len(freqMap))
	i := 0
	for key, val := range freqMap {
		freqSlice[i] = data{value: key, freq: val}
		i++
	}

	// quick select O(n) average
	freqs := quickSelect(freqSlice, k, 0, len(freqSlice)-1)
	result := make([]int, len(freqs))
	fmt.Println(freqs)
	for i := 0; i < len(freqs); i++ {
		result[i] = freqs[i].value
	}
	return result
}

type data struct {
	value int
	freq  int
}

func quickSelect(freqSlice []data, k, left, right int) []data {
	// base case
	if left == right {
		return freqSlice[:k]
	}

	// partition
	leftOfPivot := partition(freqSlice, left, right)

	if leftOfPivot == k {
		return freqSlice[:k]
	}
	if leftOfPivot > k {
		return quickSelect(freqSlice, k, left, leftOfPivot-1)
	}
	// if leftOfPivot < k
	return quickSelect(freqSlice, k, leftOfPivot+1, right)
}

func partition(freqSlice []data, left, right int) int {
	pivotIndex := left + rand.Intn(right-left)
	pivot := freqSlice[pivotIndex].freq
	freqSlice[pivotIndex], freqSlice[right] = freqSlice[right], freqSlice[pivotIndex]
	swapIndex := left
	for i := left; i <= right; i++ {
		if freqSlice[i].freq > pivot {
			freqSlice[i], freqSlice[swapIndex] = freqSlice[swapIndex], freqSlice[i]
			swapIndex++
		}
	}
	freqSlice[swapIndex], freqSlice[right] = freqSlice[right], freqSlice[swapIndex]
	return swapIndex
}
