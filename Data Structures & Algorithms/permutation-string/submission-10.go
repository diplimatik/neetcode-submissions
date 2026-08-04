func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Arr := [26]int{}
	for _, c := range s1 {
		s1Arr[byte(c)-'a']++
	}
	s2Arr := [26]int{}
	for i := 0; i < len(s1); i++ {
		s2Arr[s2[i]-'a']++
	}

	for l, r := 0, len(s1)-1; r < len(s2)-1; l++ {
		if s1Arr != s2Arr {
			r++
			s2Arr[s2[l]-'a']--
			s2Arr[s2[r]-'a']++
		} else {
			return true
		}
	}
	return s1Arr == s2Arr
}