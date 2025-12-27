package p198

func rob(nums []int) int {
	a, b := 0, 0

	for _, n := range nums {
		a, b = b, max(b, a+n)
	}

	return b
}
