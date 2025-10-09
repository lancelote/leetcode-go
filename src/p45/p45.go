package p45

func jump(nums []int) int {
	steps := 0
	curEnd := 0
	maxIdx := 0

	for i := 0; i < len(nums)-1; i++ {
		maxIdx = max(maxIdx, i+nums[i])

		if i == curEnd {
			curEnd = maxIdx
			steps++
		}
	}

	return steps
}
