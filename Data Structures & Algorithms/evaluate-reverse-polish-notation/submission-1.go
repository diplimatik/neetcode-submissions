func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, t := range tokens {
		if num, err := strconv.Atoi(t); err == nil {
			stack = append(stack, num)
		} else {
			pop1 := stack[len(stack)-1]
			pop2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch t {
			case "+":
				stack = append(stack, pop1+pop2)
			case "-":
				stack = append(stack, pop2-pop1)
			case "*":
				stack = append(stack, pop1*pop2)
			case "/":
				stack = append(stack, pop2/pop1)
			}

		}
	}
	return stack[0]
}