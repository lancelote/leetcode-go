package p724

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0

	for _, x := range nums {
		rightSum += x
	}

	for i, x := range nums {
		rightSum -= x

		if leftSum == rightSum {
			return i
		}

		leftSum += x
	}

	return -1
}
