package data_structures

func lengthOfLongestSubstring(s string) int {
	max, m := 0, map[byte]int{} // golang has no char so could either use byte or rune here
	for substringStart, substringFinish := 0, 0; substringFinish < len(s); substringFinish++ {
		// if map already has character at substringFinish increment substringStart
		if j, ok := m[s[substringFinish]]; ok && j >= substringStart {
			substringStart = j + 1 // could jump to substringStart max (k + 1) but math.Max only uses float
		}
		m[s[substringFinish]] = substringFinish
		// if new substring is longer then best so far update
		if l := substringFinish - substringStart + 1; l > max {
			max = l
		}
	}
	return max
}
