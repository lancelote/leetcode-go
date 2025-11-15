package p1502

import (
	"slices"
)

func canMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)

	step := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != step {
			return false
		}
	}

	return true
}
