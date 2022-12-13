package tutorial30daysofcode

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)
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

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
