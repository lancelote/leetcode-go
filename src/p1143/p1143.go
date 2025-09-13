package p1143

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		newDp := make([]int, len(text2)+1)

		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				newDp[j] = dp[j-1] + 1
			} else {
				newDp[j] = max(newDp[j-1], dp[j])
			}
		}

		dp = newDp
	}

	return dp[len(dp)-1]
}
