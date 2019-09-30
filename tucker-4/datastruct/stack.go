package datastruct

type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	return &Stack{ll: &LinkedList{}}
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}

func (s *Stack) Empty() bool {
	return s.Empty()
}