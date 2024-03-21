package data_structures

import "sort"

// approaches
// * intuition -> sort letters by count.. do highest count first... work way through all letters or < n -> should always do highest letter first if possible
// * heap but how representing pausing? maybe just sort once?

// specifics
// *
func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int, 0)
	for _, task := range tasks {
		counts[task]++
	}

	// sort counts
	sorted := make([]data, 0)
	for k, v := range counts {
		sorted = append(sorted, data{task: k, count: v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count // decreasing
	})
	// loop through n -> reducing counts..
	maxFreq := sorted[0].count - 1
	idleSlots := maxFreq * n
	for i := 1; i < len(sorted); i++ {
		idleSlots -= min(maxFreq, sorted[i].count) // if B has 3 and A has 4 then there is an idleSlot
	}
	if idleSlots < 0 {
		return len(tasks)
	}
	return len(tasks) + idleSlots
}

type data struct {
	task  byte
	count int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
