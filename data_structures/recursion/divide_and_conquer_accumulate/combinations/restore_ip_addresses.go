package data_structures

// clarifying questions

// approaches
// * typically these sort of questions are dynamic programming because one decision affects a later decision.. 
// * however maybe can just do a divide and conquer approach with accumulation?
// * can we use the len of s to minimize search space e.g. could put dot in the middle
import "strconv"

func restoreIpAddresses(s string) []string {
    return recurse(s, "", 0)
}

func recurse(s string, acc string, numSegments int) []string {
    // base cases
    if len(s) == 0 && numSegments == 4 {
        return []string{acc[1:]} // remove "."
    }
    if len(s) == 0 || numSegments >= 4 {
        return []string{}
    }
    // recursive cases
    // single number - always valid
    single := recurse(s[1:], acc + "." + string(s[0]), numSegments + 1)
    // double number
    var double []string
    if len(s) >= 2 {
        doubleNum, _ := strconv.Atoi(s[:2])
        if doubleNum >= 10 { // either leading 0 or not enough string len
            double = recurse(s[2:], acc + "." + s[:2], numSegments + 1)
        }
    }
    var triple []string
    if len(s) >= 3 {
        tripleNum, _ := strconv.Atoi(s[:3])
        if tripleNum >= 100 && tripleNum <= 255 {
            triple = recurse(s[3:], acc + "." + s[:3], numSegments + 1)
        }
    }
    return append(single, append(double, triple...)...)
}