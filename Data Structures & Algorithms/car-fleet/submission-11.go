func carFleet(target int, position []int, speed []int) int {
	sortedFleetIdx := make([]int, len(position))
	for i := 0; i < len(position); i++ {
		sortedFleetIdx[i] = i
	}
	sort.SliceStable(sortedFleetIdx, func(i, j int) bool {
		return position[sortedFleetIdx[i]] > position[sortedFleetIdx[j]]
	})
	leadingCarIdx := 0
	fleetCount := 1
	for i := 1; i < len(sortedFleetIdx); i++ {
		if float64(target-position[sortedFleetIdx[leadingCarIdx]])/float64(speed[sortedFleetIdx[leadingCarIdx]]) <
			float64(target-position[sortedFleetIdx[i]])/float64(speed[sortedFleetIdx[i]]) {
			leadingCarIdx = i
			fleetCount++
		}
	}
	return fleetCount
}