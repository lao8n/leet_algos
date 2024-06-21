package data_structures

import "sort"

// approaches
// * basically we want to sort the points by the minimum square that contains them - then we start from smallest until we get duplicate
func maxPointsInsideSquare(points [][]int, s string) int {
	// make point slice
	ps := make([]point, len(points))
	for i, p := range points {
		ps[i] = point{x: p[0], y: p[1], c: s[i], l: getLengthSquare(p[0], p[1])}
	}

	// custom sort by length of square ->
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].l < ps[j].l
	})
	// check for duplicates with set
	set := make(map[byte]struct{}, 0)
	maxPoints := 0
	i := 0
	for i < len(ps) {
		length := ps[i].l
		num := 0
		// check all characters with the same length
		for i < len(ps) && ps[i].l == length {
			if _, ok := set[ps[i].c]; !ok {
				num++
				set[ps[i].c] = struct{}{}
				i++
			} else {
				return maxPoints
			}
		}
		maxPoints += num
	}
	return maxPoints
}

type point struct {
	x int
	y int
	c byte
	l int
}

func getLengthSquare(x int, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x > y {
		return x
	}
	return y
}
