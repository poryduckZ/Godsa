package list

import "testing"

func TestDoublyLinkedList_Add(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)

	want := "1\n2\n3\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %s, want %s", got, want)
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)
	list.Remove(3)

	want := "1\n2\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %s, want %s", got, want)
	}
}

func TestDoublyLinkedList_RemoveHead(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)
	list.Remove(1)

	want := "2\n3\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %s, want %s", got, want)
	}
}

func TestDoublyLinkedList_RemoveNonExistent(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)
	list.Remove(6)

	want := "1\n2\n3\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %s, want %s", got, want)
	}
}
