package p290

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	byteToWord := make(map[byte]string)
	wordToByte := make(map[string]byte)

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		b := pattern[i]
		w := words[i]

		toWord, ok := byteToWord[b]
		if ok && toWord != w {
			return false
		}

		toByte, ok := wordToByte[w]
		if ok && toByte != b {
			return false
		}

		byteToWord[b] = w
		wordToByte[w] = b
	}

	return true
}
