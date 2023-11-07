package stack

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
}

func NewNode(value int) *Node {
	node := &Node{value: value}
	return node
}

func NewStack() *Stack {
	stack := &Stack{head: nil}
	return stack
}

func (s *Stack) Push(value int) *Stack {
	if s.head == nil {
		s.head = NewNode(value)
	} else {
		newNode := NewNode(value)
		newNode.next = s.head
		s.head = newNode
	}
	return s
}

func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}
	value := s.head.value
	s.head = s.head.next
	return value
}

func (s *Stack) Peek() int {
	if s.head == nil {
		return -1
	}
	return s.head.value
}
