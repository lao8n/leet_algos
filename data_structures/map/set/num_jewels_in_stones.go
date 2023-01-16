package data_structures

func numJewelsInStones(jewels string, stones string) int {
	// O(j)
	jewelsSet := make(map[byte]bool)
	for i := 0; i < len(jewels); i++ {
		jewelsSet[jewels[i]] = true
	}
	// O(s)
	numJewels := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := jewelsSet[stones[i]]; ok {
			numJewels++
		}
	}
	return numJewels
}
