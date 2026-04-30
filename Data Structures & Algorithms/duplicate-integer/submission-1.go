func hasDuplicate(nums []int) bool {
	mapToCheck := make(map[int]int)
	for _, n := range nums {
		_, ok := mapToCheck[n]
		if ok {
			return true
		}
		mapToCheck[n] = 1
	}

	return false
}