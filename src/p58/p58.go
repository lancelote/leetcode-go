package p58

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if length == 0 && s[i] == ' ' {
			continue
		}

		if s[i] == ' ' {
			return length
		}

		length++
	}

	return length
}
