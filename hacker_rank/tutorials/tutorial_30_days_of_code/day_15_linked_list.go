package tutorial_30_days_of_code

import "fmt"

type node struct {
	data int
	next *node
}

func newNode(d int) *node {
	return &node{d, nil}
}

type linkedList struct {
	head *node
}

func (l *linkedList) insert(data int) *node {
	n := newNode(data)

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
