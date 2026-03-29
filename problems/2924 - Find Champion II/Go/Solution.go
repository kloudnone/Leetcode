func findChampion(n int, edges [][]int) int {
	defeated := make([]bool, n)

	for _, edge := range edges {
		defeated[edge[1]] = true
	}

	champion := -1

	for team := 0; team < n; team++ {
		if !defeated[team] {
			if champion != -1 {
				return -1
			}
			champion = team
		}
	}

	return champion
}
