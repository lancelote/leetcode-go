package p735

func maxByAbs(a, b int) int {
	if abs(a) > abs(b) {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func willClash(a, b int) bool {
	return a > 0 && b < 0
}

func removeLast(slice []int) []int {
	return slice[:len(slice)-1]
}

func last(slice []int) int {
	return slice[len(slice)-1]
}

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

outer:
	for _, x := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, x)
			continue
		}

		for len(stack) > 0 && willClash(last(stack), x) {
			if x+last(stack) == 0 {
				stack = removeLast(stack)
				continue outer
			}

			x = maxByAbs(x, last(stack))
			stack = removeLast(stack)
		}

		stack = append(stack, x)
	}

	return stack
}
