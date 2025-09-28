package p27

func removeElement(nums []int, val int) int {
	k := 0

	for _, x := range nums {
		nums[k] = x

		if x != val {
			k++
		}
	}

	return k
}
