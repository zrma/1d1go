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
		give []interface{}
		want string
	}{
		{
			give: []interface{}{1, 2, 3},
			want: `1
2
3
`,
		},
		{
			give: []interface{}{"Hello", "World"},
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

			printArray(tt.give...)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
