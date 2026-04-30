func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	for _, ss := range s {
		sMap[ss]++
	}
	for _, tt := range t {
		val, ok := sMap[tt]
		if !ok || val == 0 {
			return false
		}
		sMap[tt]--
	}
	
	for _, val := range sMap {
		if val != 0 {
			return false
		}
	}
	
	return true
}