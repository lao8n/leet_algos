package data_structures

import "fmt"

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get func(int, int) int
 *     Dimensions func() []int
 * }
 */
// elements are 0 or 1
// each row sorted in non-decreasing order (i.e. increasing)
// which column first has a 1
// num rows or columsn <= 100

// clarifying questions: can determine anything about relative ordering of each row?
// binary search across the matrix - but we have 2 dimensions..
// could binary search within one random row

// could we sort? but don't have enough data.
// naive approach is to check every single entry i.e. m x n
// can do better where loop through every row 1 by 1 but update the columns we check everytime we find a 1...
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	m, n := dim[0], dim[1]
	leftMostColumn := n
	for i := 0; i < m; i++ { // need to check every row
		c := binarySearchRow(binaryMatrix, i, n, 0, leftMostColumn)
		if c != -1 && c < leftMostColumn {
			leftMostColumn = c
		}
	}
	if leftMostColumn == n {
		return -1
	}
	return leftMostColumn
}

func binarySearchRow(binaryMatrix BinaryMatrix, row int, numCols int, left int, right int) int {
	// recursive case
	midpoint := left + int((right-left)/2)
	fmt.Println(midpoint, left, right)
	if midpoint >= numCols {
		return -1
	}
	v := binaryMatrix.Get(row, midpoint)
	// [0, 1, 0, 0] -> left = 1, right = 2
	if left == right {
		if v == 1 {
			return midpoint
		} else {
			return -1
		}
	}

	if v == 1 { // search left
		return binarySearchRow(binaryMatrix, row, numCols, left, midpoint)
	} else if left+1 == right { // search right - i.e. side by side
		return binarySearchRow(binaryMatrix, row, numCols, right, right)
	} else {
		return binarySearchRow(binaryMatrix, row, numCols, midpoint, right)
	}
}
