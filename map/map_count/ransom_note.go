// approach
// * map of magazine with counts -> decrement for ransomeNote
func canConstruct(ransomNote string, magazine string) bool {
	counts := make(map[rune]int)
	for _, c := range magazine {
		counts[c]++
	}

	for _, c := range ransomNote {
		if counts[c] <= 0 {
			return false
		}
		counts[c]--
	}
	return true
}