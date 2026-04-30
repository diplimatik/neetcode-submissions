func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	storage := make([]map[int]struct{}, 0, len(nums))

	for i := 1; i <= len(nums)+1; i++ {
		storage = append(storage, make(map[int]struct{}))
	}

	for _, n := range nums {
		if val, ok := freqMap[n]; ok {
			delete(storage[val], n)
			freqMap[n]++
			storage[val+1][n] = struct{}{}
		} else {
			freqMap[n] = 1
			storage[1][n] = struct{}{}
		}
	}

	res := make([]int, 0, k)

	for i := len(storage) - 1; i >= 0; i-- {
		if len(storage[i]) != 0 {
			for key, _ := range storage[i] {
				if len(res) == k {
					return res
				}
				res = append(res, key)
			}
		}
	}

	return res
}