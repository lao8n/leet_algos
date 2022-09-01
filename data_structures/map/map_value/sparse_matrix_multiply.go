package data_structures

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	// brute force : given mxn and nxk matrices, complexity is m x k x n^2 = ~O(n^4), i.e.
	// each value in the final matrix has to go over n times

	// non-zero maps : we can avoid looping over the n values in a row and n values in a col
	// by storing them as maps then if on average 10% of n values are non-zero
	// complexity would be m x n + n x k + m x k x (n /10)^2 = ~O(n^2)

	// calculate cumulatively = any given entry is in final matrix is
	// r_ij = m1_i1 * m2_1j + ... m1_in * m2_nj
	// although each pair is never re-used m1_iz is re-used multiple times
	// so for m1_i1 we need m1_i1 * m2_11, m1_i1 * m2_12, ... m1_i1 * mk_1k
	// therefore we can have a map from mat2 where 1 -> [m2_11, m2_12, ... mk_1k]
	// but we only store non-zero items
	// complexity, build the mat2 map O(nxk)
	// we could build a map of mat1 - but no point as would have to loop over that
	// loop over mat1 O(mxn) and for each value we need to look up all values for that
	// column i.e. O(n), but only need to do loop for non-zero values which if we assume
	// is only 10% of the time which saves k iterations
	// O(nxk + mxn + mxn/10 * n/10) = ~O(n^2)

	// there are two ways we can build up the final mxk matrix
	// 1. calculate each mxk entry which uses mth row and kth column, although you will
	//    have to look up each component multiple times
	// 2. calculate all the values used by a component summing incrementally
	mLen := len(mat1)
	nLen := len(mat2)
	kLen := len(mat2[0])
	mat2Map := make(map[int]map[int]int) // nxk: n row -> [k col -> value]
	// create mat2 map O(nxk)
	for n := 0; n < nLen; n++ {
		// nxk matrix
		for k := 0; k < kLen; k++ {
			if mat2[n][k] != 0 {
				if map2KCol, found := mat2Map[n]; found {
					map2KCol[k] = mat2[n][k]
					mat2Map[n] = map2KCol
				} else {
					newMap2KCol := make(map[int]int)
					newMap2KCol[k] = mat2[n][k]
					mat2Map[n] = newMap2KCol
				}
			}
		}
	}
	result := make([][]int, mLen)
	// initialize matrix O(mxk)
	for m := range result {
		result[m] = make([]int, kLen)
		for k := range result[m] {
			result[m][k] = 0
		}
	}
	// add non-zero mat1[m][n] everywhere it is used O(mxnx(k/10))
	for m := 0; m < mLen; m++ {
		for n := 0; n < nLen; n++ {
			if mat1[m][n] != 0 {
				// multiply ith row of mat1 by ith col of mat2
				if mat2NRow, ok := mat2Map[n]; ok {
					for kCol, mat2NK := range mat2NRow {
						result[m][kCol] += mat1[m][n] * mat2NK
					}
				}
			}
		}
	}

	return result
}
