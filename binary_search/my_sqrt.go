package data_structures

// approach
// * use linear trial and error.. but it costs O(n)
// * user binary search trial and error - keep going until l^2 < x and r^2 >x
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		midpoint := int((l + r) / 2)
		square := midpoint * midpoint
		if square == x {
			return midpoint
		} else if square < x {
			l = midpoint + 1
		} else {
			r = midpoint - 1
		}
	}
	return r
}
