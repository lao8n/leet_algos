// q: how handle leap years? i don't know when leap years are - every year divisible by 4 is a leap year i.e. one more day for february
// q: can we go in both directions i.e. date2 > and < than date 1 - yes
// approach: split components convert to nums..

// FYI this fails one test case - apparently if the year divides by 100 it is not a leap year unless it divides by 400
func daysBetweenDates(date1 string, date2 string) int {
	y1, m1, d1 := parseDate(date1)
	y2, m2, d2 := parseDate(date2)

	// guarantee that date1 is smaller or equal to date2
	y1, m1, d1, y2, m2, d2 = orderDates(y1, m1, d1, y2, m2, d2) // guarantee that

	// sum = days up to + years between + years after
	if y1 == y2 {
		return daysUpTo(y2, m2, d2) - daysUpTo(y1, m1, d1)
	}
	return daysAfter(y1, m1, d1) + yearsBetween(y1, y2) + daysUpTo(y2, m2, d2)
}

func orderDates(y1 int, m1 int, d1 int, y2 int, m2 int, d2 int) (int, int, int, int, int, int) {
	if y1 > y2 || (y1 == y2 && m1 > m2) || (y1 == y2 && m1 == m2 && d1 > d2) {
		return y2, m2, d2, y1, m1, d1 // swap
	}
	return y1, m1, d1, y2, m2, d2
}

func parseDate(date string) (int, int, int) {
	split := strings.Split(date, "-")
	y, _ := strconv.Atoi(split[0])
	m, _ := strconv.Atoi(split[1])
	d, _ := strconv.Atoi(split[2])
	return y, m, d
}

// number days
func yearsBetween(year1 int, year2 int) int {
	if year1 > year2 { // swap to make calculation easier so year1 always bigger
		year1, year2 = year2, year1
	}
	numDays := 0
	for i := year1 + 1; i < year2; i++ {
		numDays += 365
		if i%4 == 0 {
			numDays += 1
		}
	}
	return numDays
}

// 3 scenarios - same year
// days up to, days after until
func daysUpTo(year int, month int, day int) int {
	months := map[int]int{
		2:  28, // february
		4:  30, // april
		6:  30, // june
		9:  30, // september
		11: 30, // november
	}
	numDays := 0
	for i := 1; i < month; i++ {
		if i == 2 && year%4 == 0 { // february
			numDays += 1
		}
		if d, ok := months[i]; ok {
			numDays += d
		} else {
			numDays += 31
		}
	}
	numDays += day
	return numDays
}

func daysAfter(year int, month int, day int) int {
	numDays := 365
	if year%4 == 0 {
		numDays = 366
	}
	return numDays - daysUpTo(year, month, day)
}
