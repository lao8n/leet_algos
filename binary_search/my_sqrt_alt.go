package data_structures

// approach
// * use linear trial and error.. but it costs O(n)
// * user binary search trial and error - keep going until l^2 < x and r^2 >x
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 1, x
	for l < r {
		midpoint := (l + r) / 2
		square := midpoint * midpoint
		if square == x {
			return midpoint
		} else if square < x {
			l = midpoint + 1
		} else {
			r = midpoint
		}
	}
	return l - 1
}
