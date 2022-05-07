package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestLinkedListRemoveDuplicates(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-linked-list-deletion/problem")

	for i, tt := range []struct {
		values []int
		want   []string
	}{
		{
			values: []int{1, 2, 2, 3, 3, 4},
			want:   []string{"1", "2", "3", "4"},
		},
		{
			values: []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4},
			want:   []string{"1", "2", "3", "4"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var list linkedList
			for _, v := range tt.values {
				list.head = list.insert(v)
			}

			list.head = removeDuplicates(list.head)
			got, err := utils.GetPrinted(func() {
				list.display()
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
