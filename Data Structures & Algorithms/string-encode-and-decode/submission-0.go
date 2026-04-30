type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := strconv.Itoa(len(strs)) + "*"
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "$" + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	before, encoded, _ := strings.Cut(encoded, "*")
	numOfStr, _ := strconv.Atoi(before)

	strs := make([]string, 0, numOfStr)
	for i := 0; i < numOfStr; i++ {
		before, encoded, _ = strings.Cut(encoded, "$")
		strLen, _ := strconv.Atoi(before)
		str := encoded[:strLen]
		encoded = encoded[strLen:]
		strs = append(strs, str)
	}
	return strs
}