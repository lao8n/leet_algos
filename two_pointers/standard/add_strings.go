package data_structures

import "strconv"

// approaches
// * go from right to left - easier to have spacing
// * go from left to right
func addStrings(num1 string, num2 string) string {
	i1, i2 := len(num1)-1, len(num2)-1
	carry := 0
	reversed := make([]byte, 0)
	for i1 >= 0 && i2 >= 0 {
		c1, c2 := num1[i1], num2[i2]
		d1, _ := strconv.Atoi(string(c1))
		d2, _ := strconv.Atoi(string(c2))
		num := d1 + d2 + carry
		remainder := num % 10
		carry = num / 10
		reversed = append(reversed, strconv.Itoa(remainder)[0])
		i1--
		i2--
	}
	for i1 >= 0 {
		c1 := num1[i1]
		d1, _ := strconv.Atoi(string(c1))
		num := d1 + carry
		remainder := num % 10
		carry = num / 10
		reversed = append(reversed, strconv.Itoa(remainder)[0])
		i1--
	}
	for i2 >= 0 {
		c2 := num2[i2]
		d2, _ := strconv.Atoi(string(c2))
		num := d2 + carry
		remainder := num % 10
		carry = num / 10
		reversed = append(reversed, strconv.Itoa(remainder)[0])
		i2--
	}
	if carry != 0 {
		reversed = append(reversed, strconv.Itoa(carry)[0])
	}
	reverse(reversed)
	return string(reversed)
}

func reverse(bytes []byte) {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}
}
