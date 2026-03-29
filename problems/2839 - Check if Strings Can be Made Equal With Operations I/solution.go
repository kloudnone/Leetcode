func canBeEqual(s1 string, s2 string) bool {
	if !samePair(s1[0], s1[2], s2[0], s2[2]) {
		return false
	}
	return samePair(s1[1], s1[3], s2[1], s2[3])
}

func samePair(a, b, c, d byte) bool {
	return (a == c && b == d) || (a == d && b == c)
}