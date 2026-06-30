func maxArea(heights []int) int {
	ans := 0 
	for i, j := 0, len(heights) - 1; i < j; {
		curr := (j - i) * min(heights[i], heights[j])
		if curr > ans {
			ans = curr
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}