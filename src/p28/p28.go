package p28

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0

		for j < len(needle) {
			if haystack[i+j] != needle[j] {
				break
			}

			j++
		}

		if j == len(needle) {
			return i
		}
	}

	return -1
}
