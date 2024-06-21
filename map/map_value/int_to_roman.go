package data_structures

import "bytes"

func intToRoman(num int) string {
    int_to_roman_map := map[int]string{
        1000: "M",
        500: "D",
        100: "C",
        50: "L",
        10: "X",
        5: "V",
        1: "I",
    }
    var buf bytes.Buffer
    tenOrFive := 10
    for i := 1000; i >= 1 ; {
        num = calculateRomanNumeral(tenOrFive, &buf, int_to_roman_map, num, i)
        if tenOrFive == 10 { 
            tenOrFive = 5
            i = i / 2
        } else {
            tenOrFive = 10
            i = i / 5
        }
    }
    return buf.String()
}

func calculateRomanNumeral(tenOrFive int, buf_ptr *bytes.Buffer, int_to_roman_map map[int]string, num int, multipleOfTenOrFive int) int {
    // 1000, 100, 10, or 500, 50 or 5
    for num >= multipleOfTenOrFive {
        (*buf_ptr).WriteString(int_to_roman_map[multipleOfTenOrFive])
        num -= multipleOfTenOrFive
    }
    // 900, 90, 9 or 400, 40 or 4
    if num >= multipleOfTenOrFive - multipleOfTenOrFive / tenOrFive {
        (*buf_ptr).WriteString(int_to_roman_map[multipleOfTenOrFive / tenOrFive])
        (*buf_ptr).WriteString(int_to_roman_map[multipleOfTenOrFive])
        num -= multipleOfTenOrFive - multipleOfTenOrFive / tenOrFive
    }
    return num
}

}