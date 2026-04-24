func furthestDistanceFromOrigin(moves string) int {
	var balance, blank int
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			balance--
		case 'R':
			balance++
		default:
			blank++
		}
	}
	if balance < 0 {
		return -balance + blank
	}
	return balance + blank
}
