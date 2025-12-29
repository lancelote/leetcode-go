package p139

func wordBreak(s string, wordDict []string) bool {
	var dfs func(int) bool

	cache := make(map[int]bool)

	dfs = func(pos int) bool {
		if v, ok := cache[pos]; ok {
			return v
		}

		if pos == len(s) {
			return true
		}

		r := false
		for _, w := range wordDict {
			end := pos + len(w)
			if end > len(s) {
				continue
			}

			if w != s[pos:end] {
				continue
			}

			r = dfs(pos + len(w))
			if r {
				break
			}
		}

		cache[pos] = r
		return r
	}

	return dfs(0)
}
