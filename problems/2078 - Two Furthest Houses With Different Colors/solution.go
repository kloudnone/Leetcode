func maxDistance(colors []int) int {
	n := len(colors)
	maxD := 0
	for i := 0; i < n; i++ {
		if colors[i] != colors[0] && i > maxD {
			maxD = i
		}
		if colors[i] != colors[n-1] && n-1-i > maxD {
			maxD = n - 1 - i
		}
	}

	return maxD
}