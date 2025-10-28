package p205

func isIsomorphic(s string, t string) bool {
	r := make(map[byte]byte)
	used := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		v, ok := r[s[i]]
		if !ok {
			_, exists := used[t[i]]
			if exists {
				return false
			}

			r[s[i]] = t[i]
			used[t[i]] = struct{}{}
		} else if v != t[i] {
			return false
		}
	}

	return true
}
