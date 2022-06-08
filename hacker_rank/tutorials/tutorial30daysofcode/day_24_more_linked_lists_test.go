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
		want   string
	}{
		{
			values: []int{1, 2, 2, 3, 3, 4},
			want: `1
2
3
4
`,
		},
		{
			values: []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4},
			want: `1
2
3
4
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrintln = func(a ...any) (n int, err error) {
				return fmt.Fprintln(writer, a...)
			}
			defer func() { funcPrintln = fmt.Println }()

			var list linkedList
			for _, v := range tt.values {
				list.head = list.insert(v)
			}

			list.head = removeDuplicates(list.head)
			list.display()

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
