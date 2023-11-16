package list

import (
	"testing"
)

func TestArrayList_Insert(t *testing.T) {
	l := NewArrayList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	if l.data[0] != 1 || l.data[1] != 2 || l.data[2] != 3 {
		t.Errorf("Insertion failed. Expected [1, 2, 3], got %v", l.data)
	}
}

func TestArrayList_Remove(t *testing.T) {
	l := NewArrayList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	l.Remove(2)

	if l.data[0] != 1 || l.data[1] != 3 {
		t.Errorf("Removal failed. Expected [1, 3], got %v", l.data)
	}
}

func TestArrayList_Print(t *testing.T) {
	l := NewArrayList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	expected := "[1, 2, 3]"
	result := l.Print()

	if result != expected {
		t.Errorf("Print failed. Expected %v, got %v", expected, result)
	}
}
