package list

import "testing"

func TestSinglyLinkedList_Insert(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)

	want := "1\n2\n3\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %q, want %q", got, want)
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)
	list.Remove(3)

	want := "1\n2\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %q, want %q", got, want)
	}
}

func TestSinglyLinkedList_Print(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Insert(1).Insert(2).Insert(3).Insert(4).Insert(5)

	want := "1\n2\n3\n4\n5\n"
	if got := list.Print(); got != want {
		t.Errorf("list.Print() = %q, want %q", got, want)
	}
}
