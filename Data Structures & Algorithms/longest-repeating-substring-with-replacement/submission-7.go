func characterReplacement(s string, k int) int {
	uniqueLetters := make(map[byte]int)
	l := 0
	maxLen := 0
	for r := 0; r < len(s); r++ {
		uniqueLetters[s[r]]++
		for r-l+1-uniqueLetters[s[l]] > k {
			uniqueLetters[s[l]]--
			if uniqueLetters[s[l]] == 0 {
				delete(uniqueLetters, s[l])
			}
			l++
		}

		if r == len(s)-1 && len(uniqueLetters) <= k {
			dominantLetter := s[l]
			for letter, n := range uniqueLetters {
				if n > uniqueLetters[dominantLetter] {
					dominantLetter = letter
				}
			}
			for s[l] != dominantLetter {
				l++
			}
		}

		if r == len(s)-1 && k-(r-l+1-uniqueLetters[s[l]]) > 0 {
			l -= k - (r - l + 1 - uniqueLetters[s[l]])
			if l < 0 {
				l = 0
			}
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}