package p26

func removeDuplicates(nums []int) int {
	write := 1

	for read := 1; read < len(nums); read++ {
		if nums[read-1] != nums[read] {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}
