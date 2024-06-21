package data_structures

func partitionLabels(s string) []int {
	indexMap := make(map[byte]int, 26) // map from letter to latest index
	for i := 0; i < len(s); i++ {
		indexMap[s[i]] = i
	}

	solution := make([]int, 0)
	lastIndexOfPartition, lastPartitionStart := 0, 0
	for i := 0; i < len(s); i++ {
		if indexMap[s[i]] > lastIndexOfPartition {
			lastIndexOfPartition = indexMap[s[i]]
		}
		if i == lastIndexOfPartition {
			solution = append(solution,
				lastIndexOfPartition-lastPartitionStart+1)
			lastPartitionStart = i + 1
		}
	}
	return solution
}
