package arrays

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumBribes(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/new-year-chaos/problem")

	for i, tt := range []struct {
		give []int32
		want string
	}{
		{[]int32{2, 1, 5, 3, 4}, "3"},
		{[]int32{2, 5, 1, 3, 4}, "Too chaotic"},
		{[]int32{5, 1, 2, 3, 7, 8, 6, 4}, "Too chaotic"},
		{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, "7"},
		{[]int32{1, 2, 5, 3, 4, 7, 8, 6}, "4"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			funcPrint = func(a ...any) (n int, err error) {
				return fmt.Fprint(writer, a...)
			}
			defer func() { funcPrint = fmt.Print }()

			minimumBribes(tt.give)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
