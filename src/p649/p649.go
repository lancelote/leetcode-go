package p649

func predictPartyVictory(senate string) string {
	rCount, dCount := 0, 0
	rBan, dBan := 0, 0

	q := []rune{}

	for _, senator := range senate {
		if senator == 'R' {
			rCount++
		} else {
			dCount++
		}
		q = append(q, senator)
	}

	for rCount != 0 && dCount != 0 {
		senator := q[0]
		q = q[1:]

		if senator == 'R' {
			if rBan > 0 {
				rBan--
				rCount--
			} else {
				dBan++
				q = append(q, senator)
			}
		} else {
			if dBan > 0 {
				dBan--
				dCount--
			} else {
				rBan++
				q = append(q, senator)
			}
		}
	}

	if rCount > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
