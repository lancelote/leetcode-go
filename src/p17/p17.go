package p17

var LETTERS = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	buf := make([]rune, 0, 8)

	var dfs func(int)
	dfs = func(i int) {
		if i == len(digits) {
			if len(buf) > 0 {
				result = append(result, string(buf))
			}
			return
		}

		digit := digits[i]
		for _, ch := range LETTERS[digit] {
			buf = append(buf, ch)
			dfs(i + 1)
			buf = buf[:len(buf)-1]
		}
	}

	dfs(0)
	return result
}
