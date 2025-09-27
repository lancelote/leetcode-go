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
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for i1 := 1; i1 <= len(word1); i1++ {
		for i2 := 1; i2 <= len(word2); i2++ {
			if word1[i1-1] == word2[i2-1] {
				dp[i1][i2] = dp[i1-1][i2-1]
			} else {
				dp[i1][i2] = min3(
					dp[i1-1][i2-1],
					dp[i1-1][i2],
					dp[i1][i2-1],
				) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
