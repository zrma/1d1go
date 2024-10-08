package tutorial30daysofcode

type listNode struct {
	data int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func (l *linkedList) insert(data int) *listNode {
	n := &listNode{data: data}

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

func (l *linkedList) display() {
	node := l.head
	for node != nil {
		_, _ = fmtPrintln(node.data)
		node = node.next
	}
}

func displayLinkedList(arr []int) {
	list := linkedList{nil}

	for _, data := range arr {
		list.head = list.insert(data)
	}

	list.display()
}
