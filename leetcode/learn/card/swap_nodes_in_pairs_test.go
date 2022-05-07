package card

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	t.Log("https://leetcode.com/explore/learn/card/recursion-i/250/principle-of-recursion/1681/")

	for i, tt := range []struct {
		values []int
		want   []int
	}{
		{
			values: []int{1, 2, 3, 4},
			want:   []int{2, 1, 4, 3},
		},
		{
			values: []int{1, 2, 3},
			want:   []int{2, 1, 3},
		},
		{
			values: []int{1, 2},
			want:   []int{2, 1},
		},
		{
			values: []int{1},
			want:   []int{1},
		},
		{
			values: []int{},
			want:   []int{},
		},
	} {
		t.Run(fmt.Sprintf("%d/for", i), func(t *testing.T) {
			head := buildLinkedList(tt.values)
			got := swapPairs(head)
			want := buildLinkedList(tt.want)

			gotValues := getAllValues(got)
			wantValues := getAllValues(want)
			assert.Equal(t, wantValues, gotValues)
		})

		t.Run(fmt.Sprintf("%d/recur", i), func(t *testing.T) {
			head := buildLinkedList(tt.values)
			got := swapPairsRecur(head)
			want := buildLinkedList(tt.want)

			gotValues := getAllValues(got)
			wantValues := getAllValues(want)
			assert.Equal(t, wantValues, gotValues)
		})
	}
}

func getAllValues(head *ListNode) []int {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func buildLinkedList(values []int) *ListNode {
	var head, tail *ListNode
	for _, v := range values {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}
