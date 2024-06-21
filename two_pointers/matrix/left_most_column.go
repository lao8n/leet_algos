package data_structures

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]
	i, j := 0, n-1
	for i < m && j >= 0 {
		if binaryMatrix.Get(i, j) == 0 {
			i++
		} else {
			j--
		}
	}
	if j == n-1 {
		return -1
	}
	return j + 1
}
