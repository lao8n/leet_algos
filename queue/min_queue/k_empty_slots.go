package data_structures

import (
	"fmt"
	"math"
)

// clarifying questions
// * bulbs are 0-indexed but refer to i+1th bulb

// approaches
// * after every bulb is turned on check all distances
// * after every bulb is turned on only check distances to immediate neighbours O(n^2)
// * have a map of next bulb on bulb to the right and to the left - then we add entry each way and look up left and right and distances
func kEmptySlots(bulbs []int, k int) int {
	// the day that a bulb blooms
	days := make([]int, len(bulbs))
	for day, bulb := range bulbs {
		bulb-- // 0 index bulb
		days[bulb] = day
	}

	// to have a k window be valid we need
	// [..., l, ..., r, ...] distance between l & r to be k
	// and all bloom days between l & r to be greater than l & r
	// 1. rolling window of k values = but would need to manually check if left and right are less than all of them
	// 2. minimum value = we could calculate minimum value but when that goes out of the window - have to re-calculate
	// 3. increasing values = so when first increasing value goes out we automatically have new minimum
	minQueue := []int{} // increasing list of days
	numDays := math.MaxInt
	for r, day := range days {
		// add current day to minQueue - maintaining increasing list
		for len(minQueue) > 0 && minQueue[len(minQueue)-1] > day {
			minQueue = minQueue[:len(minQueue)-1]
		}
		minQueue = append(minQueue, day)
		// min queue represents all values less than current day i
		if k <= r && r < len(days)-1 {
			// pop left to keep minQueue as a window just for next k flowers
			if days[r-k] == minQueue[0] {
				minQueue = minQueue[1:]
			}
			// [r-k]... k in between...[r + 1]
			fmt.Println(minQueue, days[r-k], days[r])
			if k == 0 || minQueue[0] > days[r-k] && minQueue[0] > days[r+1] {
				numDays = min(numDays, max(days[r-k], days[r+1]))
			}

		}
	}
	if numDays == math.MaxInt {
		return -1
	}
	return numDays + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
