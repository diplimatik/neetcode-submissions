func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	mRes := make([]int, 0, len(m))
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
		mRes = append(mRes, resNum)
	}
	resMax := 0
	for _, n := range mRes {
		if resMax < n {
			resMax = n
		}
	}
	return resMax
}