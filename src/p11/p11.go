package p11

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var maxWater int

	for left < right {
		water := min(height[left], height[right]) * (right - left)
		maxWater = max(maxWater, water)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
