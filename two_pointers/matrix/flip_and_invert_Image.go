package data_structures

// clarifying questions
// flip horizontal? i.e. around middle column
// what does invert mean?

// approaches
// * just manually do the process - flip is O(n^2/2) as swap half of elements with other half.
// * invert is O(n^2) again
// * don't think can do anything clever as processes are independent... i suppose could flip and invert at the same time? to save an O(n^2)
func flipAndInvertImage(image [][]int) [][]int {
	// flip
	n := len(image)
	// n = 3 -> n/2 = 1 so 0 and 2 but not 1
	// n = 4 -> n/2 = 2 so 0 - 1 and 3 2 -which is correct
	for i := 0; i < n; i++ {
		for j := 0; j < (n+1)/2; j++ {
			image[i][j], image[i][n-j-1] = getInverse(image[i][n-j-1]), getInverse(image[i][j])
		}
	}
	return image
}

func getInverse(x int) int {
	if x == 0 {
		return 1
	}
	return 0
}
