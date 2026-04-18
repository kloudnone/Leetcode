func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	limit := n / 2

	for d := 0; d <= limit; d++ {
		right := (startIndex + d) % n
		if words[right] == target {
			return d
		}

		if d == 0 {
			continue
		}

		left := (startIndex - d + n) % n
		if left != right && words[left] == target {
			return d
		}
	}

	return -1
}