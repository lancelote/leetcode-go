package p2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	seen1 := make(map[int]struct{})
	seen2 := make(map[int]struct{})

	for _, x := range nums1 {
		seen1[x] = struct{}{}
	}

	for _, x := range nums2 {
		seen2[x] = struct{}{}
	}

	only1 := []int{}
	only2 := []int{}

	for x := range seen1 {
		if _, ok := seen2[x]; !ok {
			only1 = append(only1, x)
		}
	}

	for x := range seen2 {
		if _, ok := seen1[x]; !ok {
			only2 = append(only2, x)
		}
	}

	return [][]int{only1, only2}
}
