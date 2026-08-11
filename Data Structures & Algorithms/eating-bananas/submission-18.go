func minEatingSpeed(piles []int, h int) int {
	pace := piles[0]
	for _, pile := range piles {
		pace = max(pile, pace)
	}
	halfPace := pace / 2
	for pace >= 1 && halfPace >= 1 {
		timeSpent := 0
		for _, pile := range piles {
			if pile%halfPace == 0 && pile >= halfPace {
				timeSpent += pile / halfPace
			} else {
				timeSpent += pile/halfPace + 1
			}
		}
		if timeSpent <= h {
			pace = halfPace
			halfPace /= 2
		} else {
			if (pace+halfPace)/2+1 == pace {
				break
			}
			halfPace = (pace+halfPace)/2 + 1
		}
	}

	if halfPace == 0 {
		halfPace = 1
	}
	for i := halfPace; i < pace; i++ {
		timeSpent := 0
		for _, pile := range piles {
			if pile%i == 0 && pile >= i {
				timeSpent += pile / i
			} else {
				timeSpent += pile/i + 1
			}
		}
		if timeSpent <= h {
			return i
		}
	}

	return pace
}