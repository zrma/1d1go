package tutorial30daysofcode

func removeDuplicates(head *listNode) *listNode {
	cur := head
	for cur != nil && cur.next != nil && cur != cur.next {
		if cur.data == cur.next.data {
			cur.next = cur.next.next
			continue
		}
		cur = cur.next
	}
	return head
}
