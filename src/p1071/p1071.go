package p1071

func isDivisibleBy(str1 string, str2 string) bool {
	for j := 0; j < len(str1)/len(str2); j++ {
		for i := 0; i < len(str2); i++ {
			shift := len(str2) * j
			if str1[i+shift] != str2[i] {
				return false
			}
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	var shortest int
	if len1 < len2 {
		shortest = len1
	} else {
		shortest = len2
	}

	for i := shortest; i > 0; i-- {
		if len1%i != 0 || len2%i != 0 {
			continue
		}
		substr := str1[0:i]
		if isDivisibleBy(str1, substr) && isDivisibleBy(str2, substr) {
			return substr
		}
	}

	return ""
}
