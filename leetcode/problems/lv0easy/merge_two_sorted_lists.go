package lv0easy

import (
	"1d1go/leetcode/problems/common"
)

func mergeTwoLists(l1, l2 *common.ListNode) *common.ListNode {
	res := &common.ListNode{}
	cur := res
	for l1 != nil || l2 != nil {

		if l1 == nil {
			cur.Next = l2
			break
		}

		if l2 == nil {
			cur.Next = l1
			break
		}

		v1 := l1.GetVal()
		v2 := l2.GetVal()
		if v1 < v2 {
			cur.Next = l1
			cur = l1
			l1 = l1.GetNext()
		} else {
			cur.Next = l2
			cur = l2
			l2 = l2.GetNext()
		}
	}
	return res.GetNext()
}
