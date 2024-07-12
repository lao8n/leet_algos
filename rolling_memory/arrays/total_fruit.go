package data_structures

func totalFruit(fruits []int) int {
	curNumIndices := make([]int, 0)
	prevNumIndices := make([]int, 0)
	maxSeq := 0
	for i, fruit := range fruits {
		if i == 0 || fruit == fruits[curNumIndices[0]] {
			curNumIndices = append(curNumIndices, i)
		} else {
			// second new number
			if len(prevNumIndices) == 0 {
				prevNumIndices = curNumIndices
				curNumIndices = []int{i}
			} else {
				// returning to prev number
				if fruit == fruits[prevNumIndices[0]] {
					curNumIndices, prevNumIndices = prevNumIndices, curNumIndices
					curNumIndices = append(curNumIndices, i)
					// entirely new number
				} else { // we want to find the earliest curNumIndex before prevNum
					lastPrevNumIndex := prevNumIndices[len(prevNumIndices)-1]
					prevNumIndices = make([]int, 0)
					for _, ci := range curNumIndices {
						if ci > lastPrevNumIndex {
							prevNumIndices = append(prevNumIndices, ci)
						}
					}
					curNumIndices = []int{i}
				}
			}
		}
		// update max seq
		maxSeq = max(maxSeq, i-curNumIndices[0]+1)
		if len(prevNumIndices) > 0 {
			maxSeq = max(maxSeq, i-prevNumIndices[0]+1)
		}
	}
	return maxSeq
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
