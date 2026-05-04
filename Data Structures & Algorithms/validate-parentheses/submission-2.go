func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		if len(stack) != 0 && (c-stack[len(stack)-1] == 1 || c-stack[len(stack)-1] == 2) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}