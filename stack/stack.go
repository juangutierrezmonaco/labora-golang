package stack

type stack []int

func NewStack() *stack {
	stack := make(stack, 0)
	return &stack
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	stack := *s
	head := stack[len(stack)-1]
	return head
}

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	stack := *s
	l := len(stack)
	head := stack[l-1]
	*s = stack[:l-1]
	return head
}