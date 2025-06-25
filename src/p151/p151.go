package p151

import (
	"strings"
)

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	reverseSlice(words)
	return strings.Join(words, " ")
}
