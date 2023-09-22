package data_structures

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	billions := num / 1000000000
	num = num % 1000000000

	millions := num / 1000000
	num = num % 1000000

	thousands := num / 1000
	num = num % 1000

	s := []string{}
	if billions > 0 {
		s = append(s, toStr(billions)...)
		s = append(s, "Billion")
	}
	if millions > 0 {
		s = append(s, toStr(millions)...)
		s = append(s, "Million")
	}
	if thousands > 0 {
		s = append(s, toStr(thousands)...)
		s = append(s, "Thousand")
	}
	s = append(s, toStr(num)...)
	return strings.Join(s, " ")
}

func toStr(n int) []string {
	s := []string{}
	digitNames := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	tenNames := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	teenNames := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	hundreds := n / 100
	n = n % 100
	tens := n / 10
	n = n % 10
	if hundreds > 0 {
		s = append(s, digitNames[hundreds], "Hundred")
	}
	if tens > 0 {
		if tens == 1 {
			s = append(s, teenNames[n])
			n = 0
		} else {
			s = append(s, tenNames[tens])
		}
	}
	if n > 0 {
		s = append(s, digitNames[n])
	}
	return s
}
