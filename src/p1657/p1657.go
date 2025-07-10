package p1657

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	count1 := make(map[rune]int)
	for _, x := range word1 {
		count1[x]++
	}

	count2 := make(map[rune]int)
	for _, x := range word2 {
		count2[x]++
	}

	if len(count1) != len(count2) {
		return false
	}

	for k, _ := range count1 {
		if _, ok := count2[k]; !ok {
			return false
		}
	}

	freq1 := make([]int, 0, len(count1))
	for _, v := range count1 {
		freq1 = append(freq1, v)
	}

	freq2 := make([]int, 0, len(count2))
	for _, v := range count2 {
		freq2 = append(freq2, v)
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
