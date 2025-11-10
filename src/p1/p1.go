package p1

func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)

	for i, num := range nums {
		idx, ok := complements[target-num]
		if ok {
			return []int{idx, i}
		}
		complements[num] = i
	}

	return []int{-1, -1}
}
