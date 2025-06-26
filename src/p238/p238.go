package p238

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))

	for i := 0; i < len(r); i++ {
		r[i] = 1
	}

	total := nums[0]

	for i := 1; i < len(nums); i++ {
		r[i] = total
		total *= nums[i]
	}

	total = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		r[i] *= total
		total *= nums[i]
	}

	return r
}
