package p2390

func removeStars(s string) string {
	stack := []rune{}

	for _, x := range s {
		if x == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, x)
		}
	}

	return string(stack)
}
