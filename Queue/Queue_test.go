package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(NewNode(1))
	queue.Enqueue(NewNode(2))
	queue.Enqueue(NewNode(3))
	queue.Enqueue(NewNode(4))
	queue.Enqueue(NewNode(5))

	want := 1
	if got := queue.Peek(); got != want {
		t.Errorf("queue.Peek() = %d, want %d", got, want)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(NewNode(1))
	queue.Enqueue(NewNode(2))
	queue.Enqueue(NewNode(3))
	queue.Enqueue(NewNode(4))
	queue.Enqueue(NewNode(5))
	queue.Dequeue()

	want := 2
	if got := queue.Peek(); got != want {
		t.Errorf("queue.Peek() = %d, want %d", got, want)
	}
}

func TestQueue_Peek(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(NewNode(1))
	queue.Enqueue(NewNode(2))
	queue.Enqueue(NewNode(3))
	queue.Enqueue(NewNode(4))
	queue.Enqueue(NewNode(5))

	want := 1
	if got := queue.Peek(); got != want {
		t.Errorf("queue.Peek() = %d, want %d", got, want)
	}
}

func TestQueue_PeekEmpty(t *testing.T) {
	queue := NewQueue()

	want := -1
	if got := queue.Peek(); got != want {
		t.Errorf("queue.Peek() = %d, want %d", got, want)
	}
}
