func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	maxSequence := 0
	for key, _ := range m {
		resMax := key
		resNum := 1
		for _, ok := m[resMax-resNum]; ok; {
			delete(m, resMax-resNum)
			resNum++
			_, ok = m[resMax-resNum]
		}
		for _, ok := m[resMax+1]; ok; {
			delete(m, resMax+1)
			resMax++
			resNum++
			_, ok = m[resMax+1]
		}
		if resNum > maxSequence {
			maxSequence = resNum
		}
	}
	return maxSequence
}