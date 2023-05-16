package data_structures

func kWeakestRows(mat [][]int, k int) []int {
	solution := make([]int, k)
	i := 0
	for c := 0; c < len(mat[0]); c++ {
		for r := 0; r < len(mat); r++ {
			if mat[r][c] == 0 && (c == 0 || mat[r][c-1] == 1) {
				solution[i] = r
				i++
				if i == k {
					return solution
				}
			}
		}
	}
	for r := 0; r < len(mat); r++ {
		if mat[r][len(mat[0])-1] == 1 {
			solution[i] = r
			i++
			if i == k {
				return solution
			}
		}
	}
	return solution
}
