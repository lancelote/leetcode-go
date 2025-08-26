package p452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	limit := points[0][1]

	for i := 1; i < len(points); i++ {
		balloon := points[i]
		start, end := balloon[0], balloon[1]

		if start <= limit {
			continue
		} else {
			arrows++
			limit = end
		}
	}

	return arrows
}
