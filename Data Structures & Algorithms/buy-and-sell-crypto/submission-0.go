func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	minIdx := -1
	i := 0
	for j := 1; j < len(prices); j++ {
		if prices[j] < prices[i] {
			i = j
		} else {
			if minIdx == -1 {
				minIdx = j - 1
				for k := j - 2; k > 0; k-- {
					if prices[k] < prices[minIdx] {
						minIdx = k
					}
				}
				i = minIdx
			}
			if profit < prices[j]-prices[i] {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}