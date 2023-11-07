package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	want := 5
	if got := stack.Pop(); got != want {
		t.Errorf("stack.Pop() = %d, want %d", got, want)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	want := 5
	if got := stack.Pop(); got != want {
		t.Errorf("stack.Pop() = %d, want %d", got, want)
	}
}

func TestStack_PopEmpty(t *testing.T) {
	stack := NewStack()

	want := -1
	if got := stack.Pop(); got != want {
		t.Errorf("stack.Pop() = %d, want %d", got, want)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	want := 5
	if got := stack.Peek(); got != want {
		t.Errorf("stack.Peek() = %d, want %d", got, want)
	}
}
