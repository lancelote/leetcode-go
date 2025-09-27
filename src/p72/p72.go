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
	var req func(int, int) int
	req = func(i1 int, i2 int) int {
		if i1 < 0 {
			return i2 + 1
		}

		if i2 < 0 {
			return i1 + 1
		}

		if word1[i1] == word2[i2] {
			return req(i1-1, i2-1)
		}

		insert := req(i1, i2-1)
		delete := req(i1-1, i2)
		replace := req(i1-1, i2-1)

		return min3(insert, delete, replace) + 1
	}

	return req(len(word1)-1, len(word2)-1)
}
