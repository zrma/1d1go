package card

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := first.Next
		pre.Next = second
		first.Next = second.Next
		second.Next = first
		pre = first
	}
	return dummy.Next
}

func swapPairsRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	first.Next = swapPairsRecur(second.Next)
	second.Next = first
	return second
}
