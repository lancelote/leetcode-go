package p72

func minDistance(word1 string, word2 string) int {
	rows, cols := len(word1)+1, len(word2)+1
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for c := 0; c < cols; c++ {
		dp[0][c] = c
	}

	for r := 0; r < rows; r++ {
		dp[r][0] = r
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			replace := dp[r-1][c-1] + 1
			delete := dp[r-1][c] + 1
			insert := dp[r][c-1] + 1

			if word1[r-1] == word2[c-1] {
				replace--
			}

			dp[r][c] = min(replace, delete, insert)
		}
	}

	return dp[rows-1][cols-1]
}
