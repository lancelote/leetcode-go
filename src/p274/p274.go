package p274

import (
	"sort"
)

func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	h := 0

	for i, x := range citations {
		h = max(h, min(x, n-i))
	}

	return h
}
