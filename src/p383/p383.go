package p383

func canConstruct(ransomNote string, magazine string) bool {
	available := make(map[rune]int)

	for _, x := range magazine {
		available[x]++
	}

	for _, x := range ransomNote {
		if available[x] == 0 {
			return false
		}

		available[x]--
	}

	return true
}
