package p2300

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	result := []int{}

	for _, s := range spells {
		left := 0
		right := len(potions) - 1
		idx := len(potions)

		for left <= right {
			middle := (left + right) / 2
			strength := int64(s * potions[middle])
			if strength >= success {
				right = middle - 1
				idx = middle
			} else {
				left = middle + 1
			}
		}

		result = append(result, len(potions)-idx)
	}

	return result
}
