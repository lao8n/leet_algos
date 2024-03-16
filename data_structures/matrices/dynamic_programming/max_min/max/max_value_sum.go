package data_structures

import "math"

// specifics
// * key thing is that if an operation is applied 2x to the same vertex then the second operation reverts the first since
// x XOR k XOR k = x
// * therefore only choice is to have one vertex XOR or not.
// * because the tree is undirected there is a single path from any vertex to any other vertex but because every vertex
// on the path (except the ends) will be x'ord twice they cancel out.
// * two number - maximum sum of elements if even number of elements were XOR ed and max sum of elements if odd number
// of elements were XORed
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	evenXored, oddXored := 0, math.MinInt
	for _, num := range nums {
		// at each step we either choose to do the xor for that num or not
		evenXored, oddXored = max(evenXored+num, oddXored+(num^k)), max(evenXored+(num^k), oddXored+num)
	}
	return int64(evenXored)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
