package p547

func findCircleNum(isConnected [][]int) int {
	result := 0
	visited := make(map[int]struct{})

	var dfs func(city int) bool
	dfs = func(city int) bool {
		if _, ok := visited[city]; ok {
			return false
		}

		visited[city] = struct{}{}

		for newCity, areConnected := range isConnected[city] {
			if areConnected == 1 {
				dfs(newCity)
			}
		}

		return true
	}

	for city := 0; city < len(isConnected); city++ {
		if dfs(city) {
			result++
		}
	}

	return result
}
