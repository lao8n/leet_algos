package data_structures

import "strings"

// n + 1 groups by n dashes
// reformat so each group contains k characters - first can be shorter but at least one
// lower case to uppercase
// approach - 1. split by dash 2. create k groups 3. format k groups (uppercase etc)
// clarifying questions? will key always be valid?
// approach - go from right to left
func licenseKeyFormatting(s string, k int) string {
	return getGroupsOfK(s, k)
}

func getGroupsOfK(s string, k int) string {
	splittedS := strings.Split(s, "-")     // [5F3Z, 2e, 9, w]
	joinedS := strings.Join(splittedS, "") //"5F3Z2e9w"
	solution := ""
	i := len(joinedS) - k
	for i > 0 {
		newK := strings.ToUpper(joinedS[i : i+k])
		solution = "-" + newK + solution
		i -= k
	}
	if i <= 0 {
		newK := strings.ToUpper(joinedS[0 : i+k])
		solution = newK + solution
	}
	return solution // 5F3z-2E9W
}
