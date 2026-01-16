package p202

func sumDigitSquares(n int) int {
	total := 0

	for n != 0 {
		d := n % 10
		n /= 10

		total += d * d
	}

	return total
}

func isHappy(n int) bool {
	seen := make(map[int]struct{})
	seen[n] = struct{}{}

	for n != 1 {
		n = sumDigitSquares(n)
		if _, exists := seen[n]; exists {
			return false
		}
		seen[n] = struct{}{}
	}

	return true
}
