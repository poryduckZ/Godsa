package stack

type StackArray struct {
	array []int
}

func NewStackArray() *StackArray {
	stackArray := &StackArray{array: []int{}}
	return stackArray
}

func (s *StackArray) Push(value int) *StackArray {
	s.array = append(s.array, value)
	return s
}

func (s *StackArray) Pop() int {
	if len(s.array) == 0 {
		return -1
	}
	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return value
}

func (s *StackArray) Peek() int {
	if len(s.array) == 0 {
		return -1
	}
	return s.array[len(s.array)-1]
}
