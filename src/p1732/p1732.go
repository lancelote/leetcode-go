package p1732

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestAltitude(gain []int) int {
	altitude := 0
	maxAltitude := 0

	for _, da := range gain {
		altitude += da
		maxAltitude = max(maxAltitude, altitude)
	}

	return maxAltitude
}
