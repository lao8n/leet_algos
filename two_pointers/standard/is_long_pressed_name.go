package data_structures

// approaches
// * two pointers - greedily match and then if get new letter in name then just skip all same letter
// * dp where track two paths one is with single match or one is not match - not necessarily

// specifics
// * loop over type of name?
// * match all of name but still have letters in typed remaining?
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		// match
		if name[i] == typed[j] {
			i++
			j++
		} else {
			// don't match
			if j > 0 && typed[j] == typed[j-1] {
				// remove long pressed
				j++
			} else {
				return false
			}
		}
	}
	// fmt.Println(i, j)
	if i < len(name) {
		return false
	}
	// remove final long pressed
	for j < len(typed) {
		if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return true
}
