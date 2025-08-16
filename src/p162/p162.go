package p162

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if middle > 0 && nums[middle-1] > nums[middle] {
			right = middle - 1
		} else if middle < len(nums)-1 && nums[middle] < nums[middle+1] {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
