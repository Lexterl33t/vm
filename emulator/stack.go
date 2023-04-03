package emulator

type Stack struct {
	Stack []int
}

func NewStack() Stack {
	st := Stack{}
	st.Stack = append(st.Stack, 0)

	return st
}

func (s *Stack) Push(value int) {
    s.Stack = append(s.Stack, value)
}

func (s *Stack) Pop() int {
    if len(s.Stack) == 0 {
        panic("la stack est vide")
    }
    value := s.Stack[len(s.Stack)-1]
    s.Stack = s.Stack[:len(s.Stack)-1]
    return value
}

func (stack *Stack) Top() int {
	if len(stack.Stack) > 0 {
		return stack.Stack[len(stack.Stack)-1]

	} else {
		stack.Stack = append(stack.Stack, 0)
		return stack.Stack[0]
	}
}
