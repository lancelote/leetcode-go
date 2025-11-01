package p242

func isAnagram(s string, t string) bool {
	sBytes := make([]int, 26)
	tBytes := make([]int, 26)

	for _, x := range s {
		i := x - 'a'
		sBytes[i]++
	}

	for _, x := range t {
		i := x - 'a'
		tBytes[i]++
	}

	for i := 0; i < 26; i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}

	return true
}
