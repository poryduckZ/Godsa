package list

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func NewNode(value int) *Node {
	node := &Node{value: value}
	return node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	singlyLinkedList := &SinglyLinkedList{head: nil}
	return singlyLinkedList
}

func (l *SinglyLinkedList) Insert(value int) *SinglyLinkedList {
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

func (l *SinglyLinkedList) Remove(value int) *SinglyLinkedList {
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

func (l *SinglyLinkedList) Print() string {
	var result string
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%d\n", current.value)
		current = current.next
	}
	return result
}
