func isGood(nums []int) bool {
	maxVal := 0
	hm := map[int]int{}
	for _, n := range nums {
		hm[n]++
		maxVal = max(maxVal, n)
	}

	for i := 1; i < maxVal; i++ {
		if hm[i] != 1 {
			return false
		}
	}

	return hm[maxVal] == 2
}