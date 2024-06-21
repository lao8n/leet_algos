package data_structures

// could start with flower then reduce for each one until "" or get to the end
// basically a stack but no need to
func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(commonPrefix) && j < len(strs[i]) {
			if !(commonPrefix[j] == strs[i][j]) {
				break
			}
			j++
		}
		commonPrefix = commonPrefix[:j]
	}
	return commonPrefix
}
