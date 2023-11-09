package queue

type Node struct {
	value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func NewNode(value int) *Node {
	node := &Node{value: value}
	return node
}

func NewQueue() *Queue {
	queue := &Queue{}
	return queue
}

func (q *Queue) Enqueue(n *Node) {
	if q.Head == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return -1
	} else {
		value := q.Head.value
		q.Head = q.Head.Next
		return value
	}
}

func (q *Queue) Peek() int {
	if q.Head == nil {
		return -1
	} else {
		return q.Head.value
	}
}
