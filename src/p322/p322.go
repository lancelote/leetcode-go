package p322

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		for x := coin; x < amount+1; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
