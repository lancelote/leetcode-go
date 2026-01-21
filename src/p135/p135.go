package p135

func candy(ratings []int) int {
	n := len(ratings)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] && result[i] <= result[i-1] {
			result[i] = result[i-1] + 1
		}
	}

	for i := n - 2; i > -1; i-- {
		if ratings[i] > ratings[i+1] && result[i] <= result[i+1] {
			result[i] = result[i+1] + 1
		}
	}

	total := 0
	for _, r := range result {
		total += r
	}

	return total
}
