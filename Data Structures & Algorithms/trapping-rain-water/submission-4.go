func trap(height []int) int {
	ans := 0
	for i, j := 0, 1; j < len(height); {
		for height[i] <= height[j] {
			if j == len(height)-1 {
				return 0
			}
			i++
			j++
		}
		maxInCase := j
		for j < len(height)-1 && height[j] < height[i] {
			j++
			if height[j] > height[maxInCase] {
				maxInCase = j
			}
		}
		if height[j] < height[i] {
			j = maxInCase
		}
		for k := i + 1; k < j; k++ {
			if height[k] >= height[i] || height[k] >= height[j] {
				continue
			}
			ans += min(height[i], height[j]) - height[k]
		}
		i = j
		j++
	}
	return ans
}