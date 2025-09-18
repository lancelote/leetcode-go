package p714

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	hold := -prices[0]
	free := 0

	for i := 1; i < len(prices); i++ {
		newHold := max(hold, free-prices[i])
		newFree := max(free, hold+prices[i]-fee)

		hold = newHold
		free = newFree
	}

	return max(hold, free)
}
