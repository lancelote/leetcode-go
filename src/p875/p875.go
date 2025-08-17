package p875

func maxItem(xs []int) int {
	if len(xs) == 0 {
		return 0
	}

	r := xs[0]
	for _, x := range xs[1:] {
		if x > r {
			r = x
		}
	}

	return r
}

func ceilDiv(a, b int) int {
	return (a + b - 1) / b
}

func willSpend(m int, piles []int) int {
	r := 0

	for _, p := range piles {
		r += ceilDiv(p, m)
	}

	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := maxItem(piles)

	minMid := right

	for left <= right {
		m := (left + right) / 2
		t := willSpend(m, piles)

		if t > h {
			left = m + 1
		} else {
			right = m - 1
			minMid = min(minMid, m)
		}
	}

	return minMid
}
