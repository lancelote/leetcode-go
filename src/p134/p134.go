package p134

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	for _, g := range gas {
		totalGas += g
	}

	totalCost := 0
	for _, c := range cost {
		totalCost += c
	}

	if totalGas < totalCost {
		return -1
	}

	total := 0
	answer := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff

		if total < 0 {
			total = 0
			answer = i + 1
		}
	}

	return answer
}
