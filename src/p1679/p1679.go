package p1679

func maxOperations(nums []int, k int) int {
	compliment := make(map[int]int)

	for _, x := range nums {
		compliment[k-x]++
	}

	result := 0
	for _, x := range nums {
		selfVal, selfOk := compliment[x]

		if selfOk && selfVal > 0 {
			compliment[x]--

			otherVal, otherOk := compliment[k-x]
			if otherOk && otherVal > 0 {
				compliment[k-x]--
				result++
			}
		}
	}

	return result
}
