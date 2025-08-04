package p399

type Pair struct {
	A string
	B string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var a, b string

	divValues := make(map[Pair]float64)
	divGraph := make(map[string][]string)

	for i, e := range equations {
		v := values[i]
		a, b = e[0], e[1]

		divValues[Pair{A: a, B: b}] = v
		divValues[Pair{A: b, B: a}] = 1 / v

		if ds, ok := divGraph[a]; ok {
			divGraph[a] = append(ds, b)
		} else {
			divGraph[a] = []string{b}
		}

		if ds, ok := divGraph[b]; ok {
			divGraph[b] = append(ds, a)
		} else {
			divGraph[b] = []string{a}
		}
	}

	visited := make(map[string]struct{})

	var dfs func(a, b string) float64
	dfs = func(a, b string) float64 {
		if _, ok := visited[a]; ok {
			return -1
		}

		visited[a] = struct{}{}

		ds, ok := divGraph[a]
		if !ok {
			delete(visited, a)
			return -1
		}

		for _, d := range ds {
			v, ok := divValues[Pair{A: a, B: d}]

			if !ok {
				continue
			}

			if d == b {
				delete(visited, a)
				return v
			}

			if r := dfs(d, b); r != -1 {
				delete(visited, a)
				return v * r
			}
		}

		delete(visited, a)
		return -1
	}

	result := []float64{}

	for _, q := range queries {
		a, b = q[0], q[1]
		result = append(result, dfs(a, b))
	}

	return result
}
