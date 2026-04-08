package main

func countValidSelections(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Calculate total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If no zeros, no valid selections
	if totalSum == 0 {
		return 0
	}

	leftSum := 0
	rightSum := totalSum
	validSelections := 0

	for _, num := range nums {
		if num != 0 {
			leftSum += num
			rightSum -= num
		} else {
			// Check if we can split the array at this position
			diff := leftSum - rightSum
			if diff == 0 {
				// Perfect split: can go left or right
				validSelections += 2
			} else if diff == 1 || diff == -1 {
				// Off by 1: only one valid direction
				validSelections += 1
			}
		}
	}

	return validSelections
}
