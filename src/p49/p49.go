package p49

func getKey(word string) [26]int {
	key := [26]int{}

	for _, x := range word {
		i := x - 'a'
		key[i]++
	}

	return key
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)

	for _, str := range strs {
		key := getKey(str)
		anagrams[key] = append(anagrams[key], str)
	}

	var result [][]string
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}
