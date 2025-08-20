package p739

type Pair struct {
	i int
	t int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []Pair{}

	for i, t := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].t < t {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last.i] = i - last.i
		}

		stack = append(stack, Pair{i: i, t: t})
	}

	return result
}
