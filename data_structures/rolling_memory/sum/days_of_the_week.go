package data_structures

// clarifying questions
// * gap years?
// * one true day? jan first 1971 is a friday

// approach?
// * convert into days from one true date and then mod 7?
func dayOfTheWeek(day int, month int, year int) string {
	days := daysFrom1971(day, month, year)
	dayOfWeek := (days + 4) % 7
	return intToDay(dayOfWeek)
}

func daysFrom1971(day int, month int, year int) int {
	fromDay, fromMonth, fromYear := 1, 1, 1971
	// gap year is every 4 years (except 100 years)
	numDays := day - fromDay
	month -= 1 // don't include partial month in full month calc
	for month >= fromMonth {
		if month == 9 || month == 4 || month == 6 || month == 11 {
			numDays += 30
		} else if month == 2 {
			numDays += 28
			if gapYear(year) {
				numDays += 1
			}
		} else {
			numDays += 31
		}
		month--
	}
	year -= 1 // don't include partial year in full year calc
	for year >= fromYear {
		numDays += 365
		if gapYear(year) {
			numDays += 1
		}
		year--
	}
	return numDays
}

func gapYear(year int) bool {
	if year%100 == 0 { // not a gap year
		if (year/100)%4 == 0 {
			return true
		}
	} else if year%4 == 0 {
		return true
	}
	return false
}

func intToDay(day int) string {
	switch day {
	case 0:
		return "Monday"
	case 1:
		return "Tuesday"
	case 2:
		return "Wednesday"
	case 3:
		return "Thursday"
	case 4:
		return "Friday"
	case 5:
		return "Saturday"
	case 6:
		return "Sunday"
	}
	return ""
}
