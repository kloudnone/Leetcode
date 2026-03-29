func removeOccurrences(s string, part string) string {
	partLen := len(part)
	for {
		index := -1
		for i := 0; i <= len(s)-partLen; i++ {
			if s[i:i+partLen] == part {
				index = i
				break
			}
		}
		if index == -1 {
			break
		}
		s = s[:index] + s[index+partLen:]
	}

	return s
}