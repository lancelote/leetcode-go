package p1318

func minFlips(a int, b int, c int) int {
	r := 0

	for a|b != c {
		ab := a & 1
		bb := b & 1
		cb := c & 1

		if ab == 1 && bb == 1 && cb == 0 {
			r += 2
		} else if ab|bb != cb {
			r += 1
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}

	return r
}
