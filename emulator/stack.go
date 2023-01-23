package emulator

type Stack struct {
	Stack []int
}

func NewStack() *Stack {
	st := Stack{}
	st.Stack = append(st.Stack, 0)

	return &st
}

func (stack *Stack) Pop() {
	stack.Stack = stack.Stack[:len(stack.Stack)-1]
}

func (stack *Stack) Push(value int) {
	stack.Stack = append(stack.Stack, value)
}

func (stack *Stack) Top() int {
	if len(stack.Stack) > 0 {
		return stack.Stack[len(stack.Stack)-1]

	} else {
		stack.Stack = append(stack.Stack, 0)
		return stack.Stack[0]
	}
}
