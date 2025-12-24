package p3074

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	totalApples := 0

	for _, pack := range apple {
		totalApples += pack
	}

	slices.Sort(capacity)

	boxUsed := 0
	for totalApples > 0 {
		totalApples -= capacity[len(capacity)-1-boxUsed]
		boxUsed++
	}

	return boxUsed
}
