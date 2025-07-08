package p1207

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	for _, x := range arr {
		count[x]++
	}

	seen := make(map[int]struct{})

	for _, v := range count {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}

	return true
}
