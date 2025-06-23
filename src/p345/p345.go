package p345

import (
	"strings"
)

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func reverseVowels(s string) string {
	vs := []rune{}

	for _, x := range s {
		if _, ok := vowels[x]; ok {
			vs = append(vs, x)
		}
	}

	var b strings.Builder
	i := len(vs) - 1

	for _, x := range s {
		if _, ok := vowels[x]; ok {
			b.WriteRune(vs[i])
			i--
		} else {
			b.WriteRune(x)
		}
	}

	return b.String()
}
