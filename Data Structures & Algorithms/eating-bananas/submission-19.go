func minEatingSpeed(piles []int, h int) int {
	maxPile := piles[0]
	for _, pile := range piles {
		maxPile = max(pile, maxPile)
	}
	
	for l, r := 1, maxPile; l <= r; {
		m := (l + r) / 2
		timeSpent := 0
		for _, pile := range piles {
			if pile%m == 0 && pile >= m {
				timeSpent += pile / m
			} else {
				timeSpent += pile/m + 1
			}
		}
		if timeSpent <= h {
			r = m-1
			maxPile=m
		} else {
			l = m + 1
		}
	}
	return maxPile
}