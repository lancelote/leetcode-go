package p374

func guessNumber(n int) int {
	left := 1
	right := n

	for {
		switch g := (right + left) / 2; guess(g) {
		case -1: // too high
			right = g - 1
		case 1: // too low
			left = g + 1
		case 0:
			return g
		}
	}
}
