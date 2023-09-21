package data_structures

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	verbose := 1
	str = strings.TrimSpace(str)
	result := 0

	if str == "" {
		return 0
	}

	if str[0] == '-' {
		verbose = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	} else if str[0] < '0' || str[0] > '9' {
		return 0
	}

	for _, j := range str {
		if j >= '0' && j <= '9' {
			result = 10*result + int(j) - '0'
			if result > math.MaxInt32 || result < math.MinInt32 {
				return math.MaxInt32*verbose + (verbose-1)/2
			}
		} else {
			break
		}
	}
	return result * verbose
}
