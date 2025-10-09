package p45

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	count := 1

	i := 0
	maxIdx := nums[0]

	for maxIdx < len(nums)-1 {
		nextMaxIdx := maxIdx
		for j := i; j <= maxIdx; j++ {
			nextMaxIdx = max(nextMaxIdx, j+nums[j])
		}
		i = maxIdx

		maxIdx = nextMaxIdx
		count++
	}

	return count
}
