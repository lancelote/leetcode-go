package p1466

func minReorder(n int, connections [][]int) int {
	outPath := make(map[int][]int)
	inPath := make(map[int][]int)

	visited := make(map[int]struct{})

	for _, c := range connections {
		a := c[0]
		b := c[1]

		if v, ok := outPath[a]; !ok {
			outPath[a] = []int{b}
		} else {
			outPath[a] = append(v, b)
		}

		if v, ok := inPath[b]; !ok {
			inPath[b] = []int{a}
		} else {
			inPath[b] = append(v, a)
		}
	}

	var dfs func(city int) int
	dfs = func(city int) int {
		visited[city] = struct{}{}

		count := 0

		for _, nextCity := range inPath[city] {
			if _, ok := visited[nextCity]; !ok {
				count += dfs(nextCity)
			}
		}

		for _, nextCity := range outPath[city] {
			if _, ok := visited[nextCity]; !ok {
				count += 1 + dfs(nextCity)
			}
		}

		return count
	}

	return dfs(0)
}
