package data_structures

import "fmt"

// approaches
// * divide number to float then convert to string
// * integer division by divide by 10 for divisor
// * detect cycles?
func fractionToDecimal(numerator int, denominator int) string {
	result := ""
	if numerator*denominator < 0 {
		numerator = numerator * -1
		result = "-"
	}
	result = fmt.Sprintf("%s%d", result, numerator/denominator)
	remainder := numerator % denominator
	if remainder != 0 {
		result += "."
	}

	mapNumberIndex := map[string]int{}
	numberFound := []int{}
	for remainder != 0 {
		k := remainder * 10 / denominator
		remainder = (remainder * 10) % denominator
		signature := fmt.Sprintf("%d%d", k, remainder)
		if repeatStarter, ok := mapNumberIndex[signature]; ok {
			for i, n := range numberFound {
				if i == repeatStarter {
					result = fmt.Sprintf("%s(%d", result, n)
					continue
				}
				result = fmt.Sprintf("%s%d", result, n)
			}
			return result + ")"
		}
		numberFound = append(numberFound, k)
		mapNumberIndex[signature] = len(numberFound) - 1
	}
	for _, n := range numberFound {
		result = fmt.Sprintf("%s%d", result, n)
	}
	return result
}
