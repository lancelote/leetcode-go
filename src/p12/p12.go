package p12

import (
	"strings"
)

func intToRoman(num int) string {
	values := []struct {
		val int
		sym string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var buf strings.Builder

	for _, v := range values {
		count := num / v.val
		num -= v.val * count
		buf.WriteString(strings.Repeat(v.sym, count))
	}

	return buf.String()
}
