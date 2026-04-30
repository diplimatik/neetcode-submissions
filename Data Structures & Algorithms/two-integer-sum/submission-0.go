func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int, len(nums))
	for i, n := range nums {
		nmap[n] = i
	}
	for i, n := range nums {
		if val, ok := nmap[target-n]; ok && val != i {
			return []int{min(val, i), max(val, i)}
		}
	}
	return []int{}
}