package data_structures

// clarifying questions
// * definition - looks the same if rotated 180 degrees - i.e flipped
// * 1 - 1 yes, 2 no, 3 no, 4 no, 5, no, 6 -> 9, 7 no, 8 yes to itself
// approaches
// * is palindrome question - except different rules - approach is the same two pointers compare
// * could use a stack until half-way but find differences faster with two pointers..
func isStrobogrammatic(num string) bool {
	l, r := 0, len(num)-1
	for l <= r {
		// 6-9 or 9-6
		if (num[l] == '6' && num[r] == '9') || (num[l] == '9' && num[r] == '6') {
			l++
			r--
			// 8 - 8 or 0 or 1
		} else if (num[l] == '8' && num[r] == '8') || (num[l] == '1' && num[r] == '1') || (num[l] == '0' && num[r] == '0') {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
