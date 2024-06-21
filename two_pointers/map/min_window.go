package data_structures

func minWindow(s string, t string) string {
	// we use the two-pointer approach + map count
	tMap := map[byte]int{}
	for _, c := range []byte(t) {
		tMap[c]++
	}
	numUniqueTRemaining := len(tMap)
	minWindowString := ""
	for l, r := 0, 0; r < len(s); r++ {
		// keep r++ until found all chars
		if numUniqueTRemaining >= 0 {
			// if s[r] is in t
			if _, ok := tMap[s[r]]; ok {
				tMap[s[r]]-- // new s[r] count remaining
				if tMap[s[r]] == 0 {
					numUniqueTRemaining-- // found all the char s[r] we need
				}
			}
		}
		for numUniqueTRemaining == 0 {
			// we have a solution - let's see if it's better
			if minWindowString == "" || r-l+1 < len(minWindowString) {
				minWindowString = s[l : r+1]
			}
			if _, ok := tMap[s[l]]; ok {
				tMap[s[l]]++ // new s[l] count remaining
				if tMap[s[l]] > 0 {
					numUniqueTRemaining++
				}
			}
			l++
		}
	}
	return minWindowString
}
