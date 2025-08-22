package p338

import (
	"strconv"
)

func count(n int) int {
	r := 0

	for _, b := range strconv.FormatInt(int64(n), 2) {
		if b == '1' {
			r++
		}
	}

	return r
}

func countBits(n int) []int {
	r := make([]int, n+1)

	for i := 0; i <= n; i++ {
		r[i] = count(i)
	}

	return r
}
