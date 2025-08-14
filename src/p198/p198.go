package p198

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	twoDoors := 0
	oneDoor := nums[0]

	for i := 1; i < len(nums); i++ {
		current := max(twoDoors+nums[i], oneDoor)
		twoDoors, oneDoor = oneDoor, current
	}

	return oneDoor
}
