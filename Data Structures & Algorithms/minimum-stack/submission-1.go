type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor() MinStack {
	minIterations := make([]int, 0, 10)
	stack := make([]int, 0, 10)
	return MinStack{
		minIterations,
		stack,
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.minStack) != 0 && ms.minStack[len(ms.minStack)-1] > val {
		ms.minStack = append(ms.minStack, val)
	} else {
		if len(ms.minStack) != 0 {
			ms.minStack = append(ms.minStack, ms.minStack[len(ms.minStack)-1])
		} else {
			ms.minStack = append(ms.minStack, val)
		}
	}
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}