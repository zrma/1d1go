package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestPrintArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	for i, tt := range []struct {
		args []interface{}
		want string
	}{
		{
			args: []interface{}{1, 2, 3},
			want: `1
2
3
`,
		},
		{
			args: []interface{}{"Hello", "World"},
			want: `Hello
World
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrintln = func(a ...interface{}) (n int, err error) {
				return fmt.Fprintln(writer, a...)
			}
			defer func() { funcPrintln = fmt.Println }()

			printArray(tt.args...)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
