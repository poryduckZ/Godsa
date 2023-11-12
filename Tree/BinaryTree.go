package tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewNode(value int) *Node {
	node := &Node{value: value}
	return node
}

func NewBinaryTree() *BinaryTree {
	binaryTree := &BinaryTree{root: nil}
	return binaryTree
}

func (t *BinaryTree) Insert(value int) *BinaryTree {
	if t.root == nil {
		t.root = NewNode(value)
	} else {
		t.root.insert(value)
	}
	return t
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = NewNode(value)
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(value)
		} else {
			n.right.insert(value)
		}
	}
}

func (t *BinaryTree) Remove(value int) *BinaryTree {
	if t.root == nil {
		return t
	}

	if t.root.value == value {
		t.root = t.root.remove()
		return t
	}

	t.root.removeNode(value)
	return t
}

func (n *Node) removeNode(value int) {
	if value < n.value {
		if n.left != nil {
			if n.left.value == value {
				n.left = n.left.remove()
				return
			}
			n.left.removeNode(value)
		}
	} else {
		if n.right != nil {
			if n.right.value == value {
				n.right = n.right.remove()
				return
			}
			n.right.removeNode(value)
		}
	}
}

func (n *Node) remove() *Node {
	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	n.value = n.right.getMinValue()
	n.right = n.right.removeMinValue()
	return n
}

func (n *Node) getMinValue() int {
	if n.left == nil {
		return n.value
	}
	return n.left.getMinValue()
}

func (n *Node) removeMinValue() *Node {
	if n.left == nil {
		return n.right
	}
	n.left = n.left.removeMinValue()
	return n
}

func (t *BinaryTree) Print() string {
	var result string
	t.root.print(&result)
	return result
}

func (n *Node) print(result *string) {
	if n.left != nil {
		n.left.print(result)
	}
	*result += fmt.Sprintf("%d\n", n.value)
	if n.right != nil {
		n.right.print(result)
	}
}
