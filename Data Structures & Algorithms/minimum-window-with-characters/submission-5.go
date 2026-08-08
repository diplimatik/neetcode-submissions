func minWindow(s string, t string) string {
	tMap := make(map[byte]int, len(t))
	for _, c := range t {
		tMap[byte(c)]++
	}
	ans := ""
	counterFound := 0
	currMap := make(map[byte]int, len(t))

	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := tMap[s[r]]; ok {
			currMap[s[r]]++
			if currMap[s[r]] <= tMap[s[r]] {
				counterFound++
			}
			if counterFound == len(t) {
				for val, ok := currMap[s[l]]; l < len(s); {
					if val > tMap[s[l]] {
						currMap[s[l]]--
					} else if ok && val == tMap[s[l]] {
						break
					}
					l++
					val, ok = currMap[s[l]]
				}
				if ans == "" || len(ans) > r-l+1 {
					ans = s[l : r+1]
				}
				currMap[s[l]]--
				l++
				counterFound--
			}
		}
	}
	return ans
}
