package p88

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1
	k := n + m - 1

	for i2 >= 0 {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[k] = nums1[i1]
			i1--
		} else {
			nums1[k] = nums2[i2]
			i2--
		}
		k--
	}
}
