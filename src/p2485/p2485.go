package p2485

func pivotInteger(n int) int {
	left := 0
	right := n * (n + 1) / 2

	for i := 1; i <= n; i++ {
		left += i

		if left == right {
			return i
		}

		right -= i
	}

	return -1
}
