package tutorial30daysofcode

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassAndInstance(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-class-vs-instance/problem")

	for _, tt := range []struct {
		give int
		want string
	}{
		{
			give: -1,
			want: `Age is not valid, setting age to 0.
You are young.
You are young.
`,
		},
		{
			give: 10,
			want: `You are young.
You are a teenager.
`,
		},
		{
			give: 16,
			want: `You are a teenager.
You are old.
`,
		},
		{
			give: 18,
			want: `You are old.
You are old.
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.give), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			funcPrintln = func(a ...any) (n int, err error) {
				return fmt.Fprintln(writer, a...)
			}
			defer func() { funcPrintln = fmt.Println }()

			classAndInstance(tt.give)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
