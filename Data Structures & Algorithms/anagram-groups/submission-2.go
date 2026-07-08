func groupAnagrams(strs []string) [][]string {
	sMap := make(map[[26]rune][]string, len(strs))
	for _, s := range strs {
		chArr := [26]rune{}
		for _, ch := range s {
			chArr[int(ch-'a')]++
		}
		sMap[chArr] = append(sMap[chArr], s)
	}
	
	res := make([][]string, 0, len(sMap))
	
	for _, s := range sMap {
		res = append(res, s)
	}
	return res
}