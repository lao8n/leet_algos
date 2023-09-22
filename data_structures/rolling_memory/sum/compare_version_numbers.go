package data_structures

func compareVersion(version1 string, version2 string) int {
	i, j, v1Rev, v2Rev := 0, 0, 0, 0
	for i < len(version1) && j < len(version2) {
		// extract version 1 revision
		i, v1Rev = extractRevision(version1, i)
		// extract version 2 revision
		j, v2Rev = extractRevision(version2, j)
		if v1Rev < v2Rev {
			return -1
		} else if v2Rev < v1Rev {
			return 1
		}
	}
	for i < len(version1) {
		i, v1Rev = extractRevision(version1, i)
		if v1Rev != 0 { // therefore it is larger than v2
			return 1
		}
	}
	for j < len(version2) {
		j, v2Rev = extractRevision(version2, j)
		if v2Rev != 0 {
			return -1
		}
	}
	return 0
}

func extractRevision(version string, s int) (int, int) {
	i, e := s, 0
	for _, v := range version[s:] {
		if v == '.' { // found end
			return i + 1, e
		} else if v == '0' && e == 0 { // skip leading 0s
			// do nothing
		} else {
			e = e*10 + int(v-'0')
		}
		i++
	}
	return i + 1, e
}
