package p125

import (
	"strings"
)

func isAlphaNum(x byte) bool {
	return (x >= 'a' && x <= 'z') || (x >= '0' && x <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		a := s[l]
		b := s[r]

		if !isAlphaNum(a) {
			l++
		} else if !isAlphaNum(b) {
			r--
		} else if a != b {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}
