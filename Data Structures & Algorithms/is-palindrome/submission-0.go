func isPalindrome(s string) bool {
	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = regex.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	for i, c := range s {
		if uint8(c) != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
