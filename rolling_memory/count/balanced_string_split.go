package data_structures

// approaches
// * dp problem? max num?
// * greedy approach - just split whenever balanced?
func balancedStringSplit(s string) int {
	count := 0
	lCount := 0
	for _, c := range s {
		if c == 'L' {
			lCount++
		} else {
			lCount--
		}
		if lCount == 0 {
			count++
		}
	}
	return count
}
