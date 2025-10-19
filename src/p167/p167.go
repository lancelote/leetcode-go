package p167

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		a := numbers[left]
		b := numbers[right]

		if a+b == target {
			return []int{left + 1, right + 1}
		} else if a+b > target {
			right--
		} else {
			left++
		}
	}

	panic("no answer found")
}
