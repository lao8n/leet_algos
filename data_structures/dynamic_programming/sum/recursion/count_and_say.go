package data_structures

import "strconv"

// clarifying questions

// approaches
// * brute force - like fibonacci just recurse backwards O(n) but each step is getting longer
// * memoize it? -> is it actually cheaper? how memoize just short string or long ones?

// implementation
// * can we memoize further?
func countAndSay(n int) string {
	return countAndSayRecursive(make(map[string]string), n)
}

func countAndSayRecursive(memoize map[string]string, n int) string {
	// base case
	if n == 1 {
		return "1"
	}
	// recursive case
	previousStr := countAndSayRecursive(memoize, n-1)
	if memoized, ok := memoize[previousStr]; ok {
		return memoized
	}
	char := previousStr[0]
	count := 1
	curStr := ""
	for i := 1; i < len(previousStr); i++ { // count number of repeated numbers
		// repeated char
		if previousStr[i] == char {
			count++
		} else {
			// process previous chars
			curStr += strconv.Itoa(count) + string(char) // 12
			// new char
			char = previousStr[i]
			count = 1
		}
	}
	curStr += strconv.Itoa(count) + string(char)
	return curStr
}
