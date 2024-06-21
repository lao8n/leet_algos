package data_structures

import "sort"

// matrix is a rectangle
// 1. can we decompose diagonal sort into a horizontal sort and a vertical sort?
// no because too many degrees of freedom e.g. [1, 5][1, 1] would get sorted to [1, 1][1, 5]
// 2. do it directly sort each diagonal - but how keep track of indices?
// 3. 2-step process 1. extract elements into array - sort them 2. input elements
func diagonalSort(mat [][]int) [][]int {
	diagonalElements := extractDiagonalElements(mat)
	for _, diagStart := range diagonalElements {
		values, coords := getDiagonalElementsAndIndices(mat, diagStart)
		sort.Ints(values)
		for i, c := range coords {
			mat[c.x][c.y] = values[i]
		}
	}
	return mat
}

// extract diagonal elements (can go across top row and across first column
func extractDiagonalElements(mat [][]int) []coord { // first column and first row
	diagonalElements := []coord{}
	for i, _ := range mat[0] { // first row
		diagonalElements = append(diagonalElements, coord{0, i})
	}
	for i, _ := range mat { // each row
		if i != 0 { // don't repeat top-left element
			diagonalElements = append(diagonalElements, coord{i, 0})
		}
	}
	return diagonalElements
}

// get values and sort them and then get coordinates and place them
func getDiagonalElementsAndIndices(mat [][]int, diagStart coord) ([]int, []coord) {
	i, j := diagStart.x, diagStart.y
	values, coords := []int{}, []coord{}
	for i < len(mat) && j < len(mat[0]) {
		values = append(values, mat[i][j])
		coords = append(coords, coord{i, j})
		i++
		j++
	}
	return values, coords
}

type coord struct {
	x int
	y int
}
