func checkInclusion(s1 string, s2 string) bool {
	s1Arr := [26]int{}
	for _, c := range s1 {
		s1Arr[byte(c)-'a']++
	}
	for l := 0; l < len(s2); l++ {
		if s1Arr[s2[l]-'a'] != 0 {
			compArray := [26]int{}
			compArray[s2[l]-'a']++
			for r := l + 1; r < l+len(s1) && r < len(s2); r++ {
				if s1Arr[s2[r]-'a'] != 0 {
					compArray[s2[r]-'a']++
				} else {
					l = r
					break
				}
			}
			if s1Arr == compArray {
				return true
			}
		}
	}
	return false
}