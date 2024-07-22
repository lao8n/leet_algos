package data_structures

// clarifying questions?
// * neg number? no

// approaches
// * need to extract digit either the front or back - easier to do the back
func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 {
		d := n % 10
		n = n / 10
		product *= d
		sum += d
	}
	return product - sum
}
