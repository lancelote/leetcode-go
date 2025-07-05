package p1456

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxVowels(s string, k int) int {
	left := 0
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	count := 0

	for i := 0; i < k; i++ {
		b := s[i]
		if _, ok := vowels[b]; ok {
			count++
		}
	}

	maxCount := count

	for i := k; i < len(s); i++ {
		if maxCount == k {
			return k
		}

		rb := s[i]
		if _, ok := vowels[rb]; ok {
			count++
		}

		lb := s[left]
		if _, ok := vowels[lb]; ok {
			count--
		}

		left++
		maxCount = max(maxCount, count)
	}

	return maxCount
}
