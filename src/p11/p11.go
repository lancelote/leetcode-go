package p11

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0

	for l < r {
		left := height[l]
		right := height[r]

		newArea := min(left, right) * (r - l)
		area = max(area, newArea)

		if left > right {
			r--
		} else {
			l++
		}
	}

	return area
}
