package data_structures

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// clarifying questions
// * guaranteed to exist? bad yes - good no

// approaches
// * linear search - problem is O(n)
// * binary search - pick middle until find for O(log n)

// specifics
// * first bad version has bad version true but n-1 false
func firstBadVersion(n int) int {
	return versions(0, n) // cannot be 0 indexed - can be n
}

func versions(good int, bad int) int {
	// base case
	if bad-good == 1 {
		return bad
	}
	// recursive case
	midpoint := good + (bad-good)/2 // 2 +
	if isBadVersion(midpoint) {     // found bad
		return versions(good, midpoint)
	} else {
		return versions(midpoint, bad)
	}
}
