package p994

type Pair struct {
	R int
	C int
}

func orangesRotting(grid [][]int) int {
	minutes := 0

	rotted := make(map[Pair]struct{})
	good := make(map[Pair]struct{})

	for r, row := range grid {
		for c, cell := range row {
			if cell == 1 {
				good[Pair{R: r, C: c}] = struct{}{}
			} else if cell == 2 {
				rotted[Pair{R: r, C: c}] = struct{}{}
			}
		}
	}

	shifts := []Pair{
		Pair{R: -1, C: 0},
		Pair{R: 0, C: +1},
		Pair{R: +1, C: 0},
		Pair{R: 0, C: -1},
	}

	for len(rotted) != 0 {
		newRotted := make(map[Pair]struct{})

		for pair := range rotted {
			for _, d := range shifts {
				nr := pair.R + d.R
				nc := pair.C + d.C
				newPair := Pair{R: nr, C: nc}

				if _, ok := good[newPair]; ok {
					delete(good, newPair)
					newRotted[newPair] = struct{}{}
				}
			}
		}

		if len(newRotted) > 0 {
			minutes++
		}

		rotted = newRotted
	}

	if len(good) == 0 {
		return minutes
	} else {
		return -1
	}
}
