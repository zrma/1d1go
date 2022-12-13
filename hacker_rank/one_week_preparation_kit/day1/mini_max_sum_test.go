package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniMaxSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/mini-max-sum/problem")
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum/problem")

	for i, tt := range []struct {
		give []int32
		want string
	}{
		{
			give: []int32{1, 2, 3, 4, 5},
			want: "10 14",
		},
		{
			give: []int32{1, 3, 5, 7, 9},
			want: "16 24",
		},
	} {
		t.Run(fmt.Sprintf(" %d", i), func(t *testing.T) {
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)
			funcPrintf = func(format string, a ...any) (n int, err error) {
				return fmt.Fprintf(writer, format, a...)
			}
			defer func() { funcPrintf = fmt.Printf }()

			miniMaxSum(tt.give)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
