package data_structures

// approach
// * two pointers in place
func reverseVowels(s string) string {
	runes := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isVowel(runes[l]) {
			l++
		}
		for l < r && !isVowel(runes[r]) {
			r--
		}
		// found both vowels - so swap
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
