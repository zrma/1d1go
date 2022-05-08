package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	t.Log("https://leetcode.com/explore/learn/card/recursion-i/251/scenario-i-recurrence-relation/2378/")

	head := &ListNode{Val: 0}
	cur := head
	for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		node := &ListNode{Val: n}
		cur.Next = node
		cur = node
	}

	head = reverseList(head)

	for _, n := range []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
		assert.NotNil(t, head)
		assert.Equal(t, n, head.Val)
		head = head.Next
	}
}

func TestReverseListNil(t *testing.T) {
	head := reverseList(nil)
	assert.Nil(t, head)
}
