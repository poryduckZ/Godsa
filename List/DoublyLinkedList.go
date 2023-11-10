package list

import "fmt"

type DoublyNode struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyNode(value int) *Node {
	node := &Node{value: value}
	return node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	doublyLinkedList := &DoublyLinkedList{head: nil}
	return doublyLinkedList
}

func (l *DoublyLinkedList) Insert(value int) *DoublyLinkedList {
	if l.head == nil {
		l.head = NewNode(value)
		l.tail = l.head
	} else {
		newNode := NewNode(value)
		l.tail.next = newNode
		l.tail = l.tail.next
	}
	return l
}

func (l *DoublyLinkedList) Remove(value int) *DoublyLinkedList {
	if l.head == nil {
		return l
	}
	if l.head.value == value {
		l.head = l.head.next
		return l
	}
	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next.value == value {
			currentNode.next = currentNode.next.next
			return l
		}
		currentNode = currentNode.next
	}
	return l
}

func (l *DoublyLinkedList) Print() string {
	var result string
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%d\n", current.value)
		current = current.next
	}
	return result
}
