package p1463

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubarray(nums []int) int {
	left := 0
	maxLen := 0
	extra := 1

	for right, x := range nums {
		if x == 0 {
			if extra == 1 {
				extra = 0
			} else {
				for left < right && nums[left] != 0 {
					left++
				}
				left++
			}
		}
		maxLen = max(maxLen, right-left+1)
	}

	return max(maxLen-1, 0)
}
