package p80

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	a, b := nums[0], nums[1]
	write := 2

	for read := 2; read < len(nums); read++ {
		if !(a == b && b == nums[read]) {
			nums[write] = nums[read]
			write += 1
		}
		a, b = b, nums[read]
	}

	return write
}
