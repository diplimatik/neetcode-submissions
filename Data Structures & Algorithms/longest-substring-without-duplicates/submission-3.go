func lengthOfLongestSubstring(s string) int {
	prevMap := make(map[byte]int)
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if prev, ok := prevMap[s[r]]; ok && prev >= l {
			l = prev + 1
		}
		if r - l + 1 > maxLen {
			maxLen = r - l + 1
		}
		prevMap[s[r]] = r
	}
	return maxLen
}