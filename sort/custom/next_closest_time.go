package data_structures

import (
	"fmt"
	"sort"
	"strconv"
)

// approaches
// * sort digits -> increment each digit from r -> l to see if can get larger time - alternative is difficult because need to generate next time
// * create all possible times and sort them - 4 ^ 4 = only 256 options actually
func nextClosestTime(time string) string {
	// extract digits
	digits := make(map[int]struct{}, 0)
	curTime := [4]int{}
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		d, _ := strconv.Atoi(string(time[i]))
		digits[d] = struct{}{}
		if i > 2 {
			curTime[i-1] = d
		} else {
			curTime[i] = d
		}
	}
	// fmt.Println(digits)
	// create all possible times
	times := make([][4]int, 0)
	times = append(times, [4]int{})
	for i := 0; i < 4; i++ { // 4 characters
		newTimes := make([][4]int, 0)
		for d, _ := range digits {
			for _, combination := range times {
				copyCombination := combination
				copyCombination[i] = d
				if i == 3 { // can now test validity
					if !validTime(copyCombination) {
						// filter out invalid times & duplicates
						continue
					}
				}
				newTimes = append(newTimes, copyCombination)
			}
		}
		times = newTimes
	}
	// pick next time
	sort.Slice(times, func(i, j int) bool {
		if times[i][0] != times[j][0] {
			return times[i][0] < times[j][0]
		}
		if times[i][1] != times[j][1] {
			return times[i][1] < times[j][1]
		}
		if times[i][2] != times[j][2] {
			return times[i][2] < times[j][2]
		}
		return times[i][3] < times[j][3]
	})

	// find next
	next := [4]int{}
	for i := 0; i < len(times); i++ {
		if times[i] == curTime {
			nI := (i + 1) % len(times)
			next = times[nI]
			break
		}
	}

	// convert to string
	return fmt.Sprintf("%v%v:%v%v", next[0], next[1], next[2], next[3])
}

func validTime(time [4]int) bool {
	if time[0] > 2 {
		return false
	}
	if time[0] == 2 && time[1] > 3 {
		return false
	}
	if time[2] > 5 {
		return false
	}
	return true
}
