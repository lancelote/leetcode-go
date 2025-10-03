package p169

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		panic("empty array")
	}

	count := 0
	candidate := nums[0]

	for _, x := range nums {
		if count == 0 {
			candidate = x
		}

		if x == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
