package tutorial_30_days_of_code

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(d int) *Node {
	return &Node{d, nil}
}

type linkedList struct {
	head *Node
}

func (l *linkedList) insert(data int) *Node {
	n := NewNode(data)

	if l.head == nil {
		return n
	}

	cur, next := l.head, l.head.next
	for next != nil {
		cur = cur.next
		next = cur.next
	}

	cur.next = n
	return l.head
}

func (l linkedList) display() {
	node := l.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func LinkedList(arr []int) {
	l := linkedList{nil}

	for _, data := range arr {
		l.head = l.insert(data)
	}

	l.display()
}
