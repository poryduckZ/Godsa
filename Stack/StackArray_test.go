package stack

import "testing"

func TestStackArray_Push(t *testing.T) {
	stack := NewStackArray()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	want := 5
	if got := stack.Peek(); got != want {
		t.Errorf("stack.Peek() = %d, want %d", got, want)
	}
}

func TestStackArray_Pop(t *testing.T) {
	stack := NewStackArray()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)
	stack.Pop()

	want := 4
	if got := stack.Peek(); got != want {
		t.Errorf("stack.Peek() = %d, want %d", got, want)
	}
}

func TestStackArray_PopEmpty(t *testing.T) {
	stack := NewStackArray()

	want := -1
	if got := stack.Pop(); got != want {
		t.Errorf("stack.Pop() = %d, want %d", got, want)
	}
}

func TestStackArray_Peek(t *testing.T) {
	stack := NewStackArray()
	stack.Push(1).Push(2).Push(3).Push(4).Push(5)

	want := 5
	if got := stack.Peek(); got != want {
		t.Errorf("stack.Peek() = %d, want %d", got, want)
	}
}

func TestStackArray_PeekEmpty(t *testing.T) {
	stack := NewStackArray()

	want := -1
	if got := stack.Peek(); got != want {
		t.Errorf("stack.Peek() = %d, want %d", got, want)
	}
}
