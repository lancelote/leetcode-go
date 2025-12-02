package p9

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var left int
	var right int

	digits := int(math.Log10(float64(x)))

	for range digits {
		right = x % 10
		divisor := int(math.Pow10(digits))
		if divisor != 0 {
			left = x / divisor
		} else {
			left = 0
		}

		if left != right {
			return false
		}

		if divisor != 0 {
			x %= divisor
		}
		x /= 10
		digits -= 2
	}

	return true
}
