package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSorting(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-sorting/problem")

	for i, tt := range []struct {
		arr  []int32
		want string
	}{
		{
			arr: []int32{1, 2, 3},
			want: `Array is sorted in 0 swaps.
First Element: 1
Last Element: 3
`,
		},
		{
			arr: []int32{3, 2, 1},
			want: `Array is sorted in 3 swaps.
First Element: 1
Last Element: 3
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrintf = func(format string, a ...any) (n int, err error) {
				return fmt.Fprintf(writer, format, a...)
			}
			defer func() { funcPrintf = fmt.Printf }()

			sorting(tt.arr)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
