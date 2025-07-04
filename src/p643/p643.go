package p643

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	maxTotal := 0

	for i := 0; i < k; i++ {
		maxTotal += nums[i]
	}

	curTotal := maxTotal
	for i := k; i < len(nums); i++ {
		curTotal += nums[i]
		curTotal -= nums[left]
		left++

		maxTotal = max(maxTotal, curTotal)
	}

	return float64(maxTotal) / float64(k)
}
