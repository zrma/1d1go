package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-sorting/problem")

	for i, tt := range []struct {
		give []int32
		want string
	}{
		{
			give: []int32{1, 2, 3},
			want: `Array is sorted in 0 swaps.
First Element: 1
Last Element: 3
`,
		},
		{
			give: []int32{3, 2, 1},
			want: `Array is sorted in 3 swaps.
First Element: 1
Last Element: 3
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			funcPrintf = func(format string, a ...any) (n int, err error) {
				return fmt.Fprintf(writer, format, a...)
			}
			defer func() { funcPrintf = fmt.Printf }()

			sorting(tt.give)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
