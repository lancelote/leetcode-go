package p189

func reverse(nums []int, s, e int) {
	l := s
	r := e

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	kn := k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, kn-1)
	reverse(nums, kn, len(nums)-1)
}
