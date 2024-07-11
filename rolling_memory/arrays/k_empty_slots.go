package data_structures

// clarifying questions
// * bulbs are 0-indexed but refer to i+1th bulb

// approaches
// * after every bulb is turned on check all distances
// * after every bulb is turned on only check distances to immediate neighbours O(n^2)
// * have a map of next bulb on bulb to the right and to the left - then we add entry each way and look up left and right and distances
func kEmptySlots(bulbs []int, k int) int {
	on := make([]int, len(bulbs))
	for day, bulb := range bulbs {
		// convert 1 index to 0 index
		bulb--
		distLeft := -1
		for i := bulb; i >= 0; i-- {
			if on[i] == 1 {
				distLeft = bulb - i - 1
				break
			}
		}
		distRight := -1
		for i := bulb; i < len(bulbs); i++ {
			if on[i] == 1 {
				distRight = i - bulb - 1
				break
			}
		}
		if distLeft == k || distRight == k {
			return day + 1
		}
		on[bulb] = 1
	}
	return -1
}
