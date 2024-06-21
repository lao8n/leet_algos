package data_structures

func intToRoman(num int) string {
    digits := []pair {
        pair{1000, "M"}, pair{900, "CM"}, pair{500, "D"}, pair{400, "CD"}, pair{100, "C"}, pair{90, "XC"}, pair{50, "L"}, pair{40, "XL"}, pair{10, "X"}, pair{9, "IX"}, pair{5, "V"}, pair{4, "IV"}, pair{1, "I"},
    }

    var roman strings.Builder
    for _, p := range digits {
        if num == 0 {
            break
        }
        count := num / p.i
        num = num % p.i 
        for i := 0; i < count; i++ {
            roman.WriteString(p.rom)
        }
    }
    return roman.String()
}

type pair struct {
    i int
    rom string
}