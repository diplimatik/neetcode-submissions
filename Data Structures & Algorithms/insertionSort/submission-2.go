func insertionSort(pairs []Pair) [][]Pair {
	if len(pairs) == 0 {
		return [][]Pair{}
	}

	solutionSteps := make([][]Pair, 0, len(pairs))
	cpy := make([]Pair, len(pairs))
	copy(cpy, pairs)
	solutionSteps = append(solutionSteps, cpy)
	for i := 1; i < len(pairs); i++ {
		if pairs[i].Key < pairs[i-1].Key {
			for j := i - 1; j >= 0; j-- {
				if pairs[j+1].Key < pairs[j].Key {
					pairs[j+1], pairs[j] = pairs[j], pairs[j+1]
				} else {
					break
				}
			}
		}
		cpy := make([]Pair, len(pairs))
		copy(cpy, pairs)
		solutionSteps = append(solutionSteps, cpy)
	}
	return solutionSteps
}