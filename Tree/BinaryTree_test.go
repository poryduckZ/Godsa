package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if tree.root.value != 5 {
		t.Errorf("Root value should be 5, got %d", tree.root.value)
	}

	if tree.root.left.value != 3 {
		t.Errorf("Left child value should be 3, got %d", tree.root.left.value)
	}

	if tree.root.right.value != 7 {
		t.Errorf("Right child value should be 7, got %d", tree.root.right.value)
	}
}

func TestRemove(t *testing.T) {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Remove(3)

	if tree.root.left != nil {
		t.Errorf("Left child should be nil, got %d", tree.root.left.value)
	}
}
