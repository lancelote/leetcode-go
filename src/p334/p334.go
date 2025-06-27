package p334

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	a := math.MaxInt
	b := math.MaxInt

	for _, x := range nums {
		if x <= a {
			a = x
		} else if x <= b {
			b = x
		} else {
			return true
		}
	}
	return false
}
