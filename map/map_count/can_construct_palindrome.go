package data_structures

// clarifying questions
// * does it need to be a palindrome
// * does it need to be a continuous substring?

// approaches
// * dynamic programming problem - but track number of ways?

// specifics
// - have a rolling map where count numbers of letters need all to be even except one
// - l & r where build a map and slowly add r
// - might be an optimisation where we can jump to the new map but we can do that later
func canConstruct(s string, k int) bool {
	rollingMap := make(map[byte]int)
	oddCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		rollingMap[c]++
		if rollingMap[c]%2 == 0 {
			oddCount--
		} else {
			oddCount++
		}
	}
	if len(s) < k { // need at least one letter per partition
		return false
	}
	return oddCount <= k // at most one odd letter p
}
