import "strings"

func intToRoman(num int) string {
	// values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// var result strings.Builder

	// for index, value := range values {
	// 	for num >= value {
	// 		num -= value
	// 		result.WriteString(symbols[index])
	// 	}
	// }
	// return result.String()

	int_to_roman_map := make(map[int]string {
        1000: "M",
        500: "D",
        100: "C",
        50: "L"
        10: "X",
        5: "V",
        1: "I"
    })
}