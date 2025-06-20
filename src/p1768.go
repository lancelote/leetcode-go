package p1768

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var b strings.Builder
	var i1, i2 int

	for i1 < len(word1) && i2 < len(word2) {
		x := word1[i1]
		y := word2[i2]

		if i1 <= i2 {
			b.WriteByte(x)
			i1++
		} else {
			b.WriteByte(y)
			i2++
		}
	}

	for i1 < len(word1) {
		x := word1[i1]
		b.WriteByte(x)
		i1++
	}

	for i2 < len(word2) {
		y := word2[i2]
		b.WriteByte(y)
		i2++
	}

	return b.String()
}
