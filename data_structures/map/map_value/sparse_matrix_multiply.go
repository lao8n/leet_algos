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
}