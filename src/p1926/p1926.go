package p1926

type Pair struct {
	R int
	C int
}

func pairFromSlice(s []int) Pair {
	r, c := s[0], s[1]
	return Pair{R: r, C: c}
}

func nearestExit(maze [][]byte, entrance []int) int {
	shifts := []Pair{
		Pair{R: -1, C: 0},
		Pair{R: 0, C: +1},
		Pair{R: +1, C: 0},
		Pair{R: 0, C: -1},
	}
	rows, cols := len(maze), len(maze[0])

	steps := 1
	toVisit := []Pair{pairFromSlice(entrance)}
	visited := make(map[Pair]struct{})
	visited[pairFromSlice(entrance)] = struct{}{}

	for len(toVisit) > 0 {
		nextToVisit := []Pair{}

		for _, pos := range toVisit {
			for _, d := range shifts {
				nr := pos.R + d.R
				nc := pos.C + d.C

				pair := Pair{R: nr, C: nc}

				if _, ok := visited[pair]; ok {
					continue
				}

				visited[pair] = struct{}{}

				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}

				if maze[nr][nc] == '+' {
					continue
				}

				if nr == 0 || nc == 0 || nr == rows-1 || nc == cols-1 {
					return steps
				}

				nextToVisit = append(nextToVisit, pair)
			}
		}

		toVisit = nextToVisit
		steps++
	}

	return -1
}
