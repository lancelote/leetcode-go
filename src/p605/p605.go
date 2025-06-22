package p605

func canPlaceFlowers(flowerbed []int, n int) bool {
	var i int
	var left int = n

	for i < len(flowerbed) && left > 0 {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				left -= 1
				i += 2
			} else {
				i += 1
			}
		} else {
			i += 2
		}
	}

	return left == 0
}
