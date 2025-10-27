package p15

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)

	twoSum := func(a, l, r int) {
		for l < r {
			b := nums[l]
			c := nums[r]

			if a+b+c == 0 {
				result = append(result, []int{a, b, c})
				l++
				r--

				for l < r && nums[l-1] == nums[l] {
					l++
				}
			} else if a+b+c > 0 {
				r--
			} else {
				l++
			}
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		twoSum(nums[i], i+1, len(nums)-1)
	}

	return result
}
