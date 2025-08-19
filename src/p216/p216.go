package p216

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	curSum := 0
	curPath := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if curSum == n && len(curPath) == k {
			newPath := make([]int, len(curPath))
			copy(newPath, curPath)
			result = append(result, newPath)
			return
		}

		if curSum >= n {
			return
		}

		if len(curPath) > k {
			return
		}

		for j := i + 1; j <= 9; j++ {
			curSum += j
			curPath = append(curPath, j)

			dfs(j)

			curSum -= j
			curPath = curPath[:len(curPath)-1]
		}
	}

	dfs(0)
	return result
}
