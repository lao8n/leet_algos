package data_structures

// approaches
// * reverse every 2k characters
// * reverse in place using O(n)
func reverseStr(s string, k int) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i += 2 * k {
		reverse(bs, i, i+k)
	}
	return string(bs)
}

func reverse(s []byte, l, r int) { // let r be exclusive
	if r > len(s) {
		r = len(s)
	}
	length := r - l // 6 - 4
	for i := 0; i < length/2; i++ {
		sl, sr := l+i, r-i-1 // inclusive sr
		s[sl], s[sr] = s[sr], s[sl]
	}
}
