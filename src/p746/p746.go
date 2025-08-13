package p746

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	oneStep := 0
	twoSteps := 0

	for i := 2; i <= len(cost); i++ {
		minCost := min(oneStep+cost[i-1], twoSteps+cost[i-2])
		oneStep, twoSteps = minCost, oneStep
	}

	return oneStep
}
