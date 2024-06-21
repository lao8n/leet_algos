package data_structures

// check each condition manually by looping through but we have horizontal, vertical and box so that is
// O(3n), given these are partially filled, can save effort with storing all elements in a map, and
// checking them but still need to for loop O(3n) ways (where n = 9^2)
// how to check validity? again use a map otherwise need to keep for looping over list to check no repetitions
// can we memoize some of the effort? rather than checking each valid condition, what if we take each element and check validity for all 3?
// complexity if n = 9
// O(3n) to make maps
// O(n^2) to loop over all values
func isValidSudoku(board [][]byte) bool {
	horizontalValid := make([]map[byte]bool, 9)
	verticalValid := make([]map[byte]bool, 9)
	boxValid := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		horizontalValid[i] = make(map[byte]bool)
		verticalValid[i] = make(map[byte]bool)
		boxValid[i] = make(map[byte]bool)
	}
	for i, row := range board {
		for j, v := range row {
			if v != '.' {
				// 3 (row 4) 5 (row 6)
				iIndex := i / 3 // 1
				jIndex := j / 3 // 5/3 = 1
				boxValidIndex := iIndex + 3*jIndex
				if !checkValid(horizontalValid[i], v) ||
					!checkValid(verticalValid[j], v) ||
					!checkValid(boxValid[boxValidIndex], v) {
					return false
				}
			}
		}
	}
	return true
}

func checkValid(validMap map[byte]bool, v byte) bool {
	if _, ok := validMap[v]; ok {
		// already exists therefore invalid
		return false
	}
	validMap[v] = true
	return true
}
