func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	var diff [][2]byte
	for i := range s1 {
		if s1[i] != s2[i] {
			if len(diff) >= 2 {
				return false
			}
			diff = append(diff, [2]byte{s1[i], s2[i]})
		}
	}
	return len(diff) == 2 && diff[0][0] == diff[1][1] && diff[0][1] == diff[1][0]
}