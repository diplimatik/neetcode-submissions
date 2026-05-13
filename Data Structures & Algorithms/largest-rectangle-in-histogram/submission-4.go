func largestRectangleArea(heights []int) int {
	leftGreaterSlice := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		leftGreater := 0
		for j := i - 1; j >= 0; j -= leftGreaterSlice[j] + 1 {
			if heights[i] > heights[j] {
				break
			}
			leftGreater += leftGreaterSlice[j] + 1
		}
		leftGreaterSlice[i] = leftGreater
	}
	rightGreaterSlice := make([]int, len(heights))
	maxArea := 0
	for i := len(heights) - 1; i >= 0; i-- {
		rightGreater := 0
		for j := i + 1; j < len(heights); j += rightGreaterSlice[j] + 1 {
			if heights[i] > heights[j] {
				break
			}
			rightGreater += rightGreaterSlice[j] + 1
		}
		rightGreaterSlice[i] = rightGreater
		if maxArea < (rightGreaterSlice[i]+leftGreaterSlice[i]+1)*heights[i] {
			maxArea = (rightGreaterSlice[i] + leftGreaterSlice[i] + 1) * heights[i]
		}
	}
	return maxArea
}