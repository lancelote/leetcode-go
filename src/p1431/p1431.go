package p1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, x := range candies {
		if x > max {
			max = x
		}
	}

	result := make([]bool, len(candies))
	for i, x := range candies {
		if x+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
