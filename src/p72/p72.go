package p72

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var req func(int, int) int
	req = func(i1 int, i2 int) int {
		if i1 < 0 {
			return i2 + 1
		}

		if i2 < 0 {
			return i1 + 1
		}

		if memo[i1][i2] != -1 {
			return memo[i1][i2]
		}

		var minDist int
		if word1[i1] == word2[i2] {
			minDist = req(i1-1, i2-1)
		} else {
			insert := req(i1, i2-1)
			delete := req(i1-1, i2)
			replace := req(i1-1, i2-1)
			minDist = min3(insert, delete, replace) + 1
		}

		memo[i1][i2] = minDist
		return minDist
	}

	return req(len(word1)-1, len(word2)-1)
}
