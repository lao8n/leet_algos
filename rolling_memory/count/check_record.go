package data_structures

// approach
// * rolling count of absence and late O(n)
func checkRecord(s string) bool {
	absentCount, lateCount := 0, 0
	for _, c := range s {
		switch c {
		case 'A':
			absentCount++
			lateCount = 0
		case 'L':
			lateCount++
		case 'P':
			lateCount = 0
		}
		if absentCount >= 2 {
			return false
		}
		if lateCount >= 3 {
			return false
		}
	}
	return true
}
