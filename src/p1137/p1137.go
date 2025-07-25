package p1137

func tribonacci(n int) int {
	a, b, c := 0, 1, 1

	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else if n == 2 {
		return c
	}

	for n > 2 {
		a, b, c = b, c, a+b+c
		n--
	}

	return c
}
