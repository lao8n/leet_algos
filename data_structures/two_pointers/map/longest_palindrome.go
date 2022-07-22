package data_structures

import (
	"fmt"
	"math"
)

// manacher's algorithm
func longestPalindrome(s string) string {
	palindromeRadii := make([]float64, 2*len(s)+1)
	midpoint, mirror, radius, mirrorRadius := 0.0, 0.0, 0.0, 0.0
	left, right := 0, 0
	for int(midpoint) < len(s) {
		// go opposite way around
		left, right = int(math.Ceil(midpoint)), int(math.Floor(midpoint))
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
			radius = float64(right) - midpoint
		}
		palindromeRadii[int(2*midpoint)] = radius
		fmt.Println(palindromeRadii)
		// mirrored radius
		mirror = midpoint
		mirrorRadius = radius
		fmt.Printf("left %v right %v midpoint %v mirror %v, mirrorRadius %v\n", left, right, midpoint, mirror, mirrorRadius)
		// reset values
		midpoint = midpoint + 0.5
		radius = 0
		for midpoint <= mirror+mirrorRadius {
			mirroredMidpoint := mirror - (midpoint - mirror)
			maxMirroredRadius := mirrorRadius + mirror - midpoint
			if palindromeRadii[int(2*mirroredMidpoint)] < maxMirroredRadius {
				palindromeRadii[int(2*midpoint)] = palindromeRadii[int(2*mirroredMidpoint)]
				midpoint++
			} else if palindromeRadii[int(2*mirroredMidpoint)] > maxMirroredRadius {
				palindromeRadii[int(2*midpoint)] = maxMirroredRadius // couldn't this not be the max
				midpoint++
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}
	// for loop
	longestLeft, longestRight, biggestRadii := 0, 0, 0.0
	for i, r := range palindromeRadii {
		if r > biggestRadii {
			biggestRadii = r
			midpoint = float64(i) / 2
			longestLeft, longestRight = int(midpoint-r), int(midpoint+r)
		}
	}
	return s[longestLeft : longestRight+1]
}
