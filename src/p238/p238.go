package p238

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))
	for i := range r {
		r[i] = 1
	}

	m := nums[0]
	for i := 1; i < len(nums); i++ {
		r[i] *= m
		m *= nums[i]
	}

	m = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] *= m
		m *= nums[i]
	}

	return r
}
