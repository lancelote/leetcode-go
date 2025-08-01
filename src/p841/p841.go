package p841

func canVisitAllRooms(rooms [][]int) bool {
	toVisit := []int{0}
	visited := make(map[int]struct{})

	for len(toVisit) > 0 {
		room := toVisit[0]
		toVisit = toVisit[1:]

		if _, ok := visited[room]; ok {
			continue // already visited
		}

		visited[room] = struct{}{}

		for _, newRoom := range rooms[room] {
			toVisit = append(toVisit, newRoom)
		}
	}

	return len(visited) == len(rooms)
}
