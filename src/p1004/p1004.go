package p1004

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0

	extra := k

	maxLen := 0

	for _, x := range nums {
		if x == 0 {
			if extra > 0 {
				extra -= 1
			} else {
				for nums[left] != 0 {
					left++
				}
				left++
			}
		}

		right++
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
