package p283

func moveZeroes(nums []int) {
	var s, f int

	for f < len(nums) && s < len(nums) {
		if nums[f] == 0 {
			f++
		} else if s < f && nums[s] == 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
			f++
		} else if s < f {
			s++
		} else {
			f++
		}
	}
}
