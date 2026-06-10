func insertionSort(pairs []Pair) [][]Pair {
	if len(pairs) == 0 {
		return [][]Pair{}
	}

	solutionSteps := make([][]Pair, 0, len(pairs))
	cpy := make([]Pair, len(pairs))
	copy(cpy, pairs)
	solutionSteps = append(solutionSteps, cpy)
	for i := 1; i < len(pairs); i++ {
		for j := i - 1; j >= 0 && pairs[j+1].Key < pairs[j].Key; j-- {
			pairs[j+1], pairs[j] = pairs[j], pairs[j+1]
		}
		cpy := make([]Pair, len(pairs))
		copy(cpy, pairs)
		solutionSteps = append(solutionSteps, cpy)
	}

	return solutionSteps
}