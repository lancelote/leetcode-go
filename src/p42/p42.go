package p42

func trap(height []int) int {
	total := 0

	l := 0
	r := len(height) - 1

	maxLeft := height[l]
	maxRight := height[r]

	for l < r-1 {
		diff := 0

		if maxLeft <= maxRight {
			l++
			left := height[l]
			diff = maxLeft - left
			maxLeft = max(maxLeft, left)
		} else {
			r--
			right := height[r]
			diff = maxRight - right
			maxRight = max(maxRight, right)
		}

		if diff > 0 {
			total += diff
		}
	}

	return total
}
