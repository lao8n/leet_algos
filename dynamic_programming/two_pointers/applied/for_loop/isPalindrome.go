package main

import (
	"fmt"
	"sort"
	"unicode"
)

type stringPalindrome []string

func (p stringPalindrome) Len() int           { return len(p) }
func (p stringPalindrome) Less(i, j int) bool { return p[i] < p[j] }
func (p stringPalindrome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// uses sort interface
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func IsPalindromeSimpler(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	var alphanumerics []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			alphanumerics = append(alphanumerics, unicode.ToLower(r))
		}
	}
	for i := range alphanumerics {
		if alphanumerics[i] != alphanumerics[len(alphanumerics)-1-i] {
			return false
		}
	}
	return true
}

// func isAlphaNumeric(a rune) bool {
//     if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9'){
//         return true
//     }
//     return false
// }

func main() {
	pal1 := []string{"a", "b", "a"}
	pal2 := []string{"a", "b", "b", "a"}
	pal3 := []string{"a", "b", "c"}
	fmt.Printf("%v is a palindrome: %v\n", pal1, IsPalindrome(stringPalindrome(pal1)))
	fmt.Printf("%v is a palindrome: %v\n", pal2, IsPalindrome(stringPalindrome(pal2)))
	fmt.Printf("%v is a palindrome: %v\n", pal3, IsPalindrome(stringPalindrome(pal3)))
}
