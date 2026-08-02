func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	prevMap := make(map[byte]int)
	maxLen := 0
	currLen := 1
	l := 0
	prevMap[s[l]] = l
	for r := 1; r < len(s); r++ {
		if prev, ok := prevMap[s[r]]; ok && prev >= l {
			l = prev + 1
			currLen = r - l + 1
		} else { // if ok && prev < l or if !ok
			currLen++
		}
		if currLen > maxLen {
			maxLen = currLen
		}
		prevMap[s[r]] = r
	}
	return maxLen
}