type histogram struct {
	leftGreater  int
	rightGreater int
	maxArea      int
}

func largestRectangleArea(heights []int) int {
	histograms := make([]histogram, len(heights))
	for i := 0; i < len(heights); i++ {
		leftGreater := 0
		for j := i - 1; j >= 0; j -= histograms[j].leftGreater + 1 {
			if heights[i] > heights[j] {
				break
			}
			leftGreater += histograms[j].leftGreater + 1
		}
		histograms[i].leftGreater = leftGreater
	}

	for i := len(heights) - 1; i >= 0; i-- {
		rightGreater := 0
		for j := i + 1; j < len(heights); j += histograms[j].rightGreater + 1 {
			if heights[i] > heights[j] {
				break
			}
			rightGreater += histograms[j].rightGreater + 1
		}
		histograms[i].rightGreater = rightGreater
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		histograms[i].maxArea = (histograms[i].leftGreater + histograms[i].rightGreater + 1) * heights[i]
		if (histograms[i].leftGreater+histograms[i].rightGreater+1)*heights[i] > maxArea {
			maxArea = (histograms[i].leftGreater + histograms[i].rightGreater + 1) * heights[i]
		}
	}
	fmt.Println(histograms)
	return maxArea
}