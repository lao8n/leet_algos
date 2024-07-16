package data_structures

// clarifying questions
// * 0, 1, 8 -> themselves
// * 2 <-> 5 and 6 <-> 9
// * 3, 4, 7 -> do not rotate

// approaches
// * number of digits -> just manually check validity of each number in range 1 -> n
// * use some rules -> to be valid new number need at least one of 2, 5, 6 or 9 and everything else cannot be 3, 4, 7
// * recursive approach break n (up to 1000) by powers of 10 ->
func rotatedDigits(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if isNumValid(i) == Good {
			count++
		}
	}
	return count
}

const (
	Good = iota
	Valid
	Invalid
)

func isNumValid(n int) int {
	state := Valid
	for n > 0 {
		digit := n % 10
		switch digitState(digit) {
		case Good:
			state = Good
		case Invalid:
			return Invalid // single invalid then bail
		}
		n /= 10
	}
	return state
}

func digitState(d int) int {
	switch d {
	case 0, 1, 8:
		return Valid
	case 2, 5, 6, 9:
		return Good
	default:
		return Invalid
	}
}
