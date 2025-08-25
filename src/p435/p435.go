package p435

import (
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	result := 0
	lastEnd := math.MinInt

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		if start >= lastEnd {
			lastEnd = end
		} else {
			result += 1
		}
	}

	return result
}
