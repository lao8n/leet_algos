package data_structures

import (
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	invalid := []string{}
	ts := make(map[string][]data) // name -> data points

	for i, s := range transactions {
		// deconstruct string
		d := strings.Split(s, ",")
		name := d[0]
		time, _ := strconv.Atoi(d[1])
		amount, _ := strconv.Atoi(d[2])
		city := d[3]

		// test validity
		invalidFlag := false
		// amount validity
		if amount > 1000 {
			invalidFlag = true
		}
		// time and city validity
		for j := range ts[name] {
			if abs(ts[name][j].time, time) <= 60 && ts[name][j].city != city {
				invalidFlag = true
				// add other invalid
				if !ts[name][j].invalid {
					otherTransaction := ts[name][j].idx
					invalid = append(invalid, transactions[otherTransaction])
					ts[name][j].invalid = true
				}
			}
		}

		// add transaction
		ts[name] = append(ts[name], data{
			time:    time,
			city:    city,
			idx:     i,
			invalid: invalidFlag,
		})

		// add if invalid
		if invalidFlag {
			invalid = append(invalid, transactions[i])
		}
	}
	return invalid
}

type data struct {
	time    int
	city    string
	idx     int
	invalid bool
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
