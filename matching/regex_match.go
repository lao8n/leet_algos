package data_structures

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	if len(p) > 1 && p[1] == '*' {
		return matchStar(s, p[2:], p[0])
	}

	if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
		return isMatch(s[1:], p[1:])
	}
	return false
}

func matchStar(s, p string, prev byte) bool {
	for {
		if isMatch(s, p) {
			return true
		}
		if len(s) == 0 || prev != s[0] && prev != '.' {
			return false
		}
		s = s[1:]
	}
}
