package data_structures

// approaches
// * brute force - try every partition with every other partition
// * memoize - partitions - but dont' need just one so need to brute force try everything
func partition(s string) [][]string {
	d := data{
		s:      s,
		result: [][]string{},
	}
	d.recurse(0, []string{})
	return d.result
}

type data struct {
	s      string
	result [][]string
}

// accumulate partitions forward
func (d *data) recurse(p int, acc []string) {
	// base cases
	if p == len(d.s) { // partition whole way to end of string
		r := make([]string, len(acc))
		copy(r, acc)
		d.result = append(d.result, r)
	}

	// recursive cases
	for i := p + 1; i <= len(d.s); i++ {
		if isPalindrome(d.s[p:i]) {
			acc = append(acc, d.s[p:i])
			d.recurse(i, acc)
			acc = acc[:len(acc)-1] // backtrack
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
