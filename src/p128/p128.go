package p128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0

	for n := range set {
		length := 1

		// start of the sequence
		if _, exists := set[n-1]; exists {
			continue
		}

		for {
			if _, exists := set[n+1]; !exists {
				break
			}

			n++
			length++
		}

		longest = max(longest, length)
	}

	return longest
}
