package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	for i, tt := range []struct {
		give []any
		want string
	}{
		{
			give: []any{1, 2, 3},
			want: `1
2
3
`,
		},
		{
			give: []any{"Hello", "World"},
			want: `Hello
World
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			fmtPrintln = func(a ...any) (n int, err error) {
				return fmt.Fprintln(writer, a...)
			}
			defer func() { fmtPrintln = fmt.Println }()

			printArray(tt.give...)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
