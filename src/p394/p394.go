package p394

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}

	curStr := ""
	num := 0

	for _, b := range s {
		if unicode.IsDigit(b) {
			digit, _ := strconv.Atoi(string(b))
			num = num*10 + digit
		} else if b == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, curStr)

			curStr = ""
			num = 0
		} else if b == ']' {
			prevNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			curStr = prevStr + strings.Repeat(curStr, prevNum)
		} else {
			curStr += string(b)
		}
	}

	return curStr
}
