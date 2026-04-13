func getMinDistance(nums []int, target int, start int) int {
	for d := 0; d < len(nums); d++ {
		if start-d >= 0 && nums[start-d] == target {
			return d
		}
		if d != 0 && start+d < len(nums) && nums[start+d] == target {
			return d
		}
	}
	return 0
}
