func maximumCount(nums []int) int {
	posCount := 0
	negCount := 0
	for _, v := range nums {
		if v > 0 {
			posCount++
		} else if v < 0 {
			negCount++
		}
	}

	return max(posCount, negCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}