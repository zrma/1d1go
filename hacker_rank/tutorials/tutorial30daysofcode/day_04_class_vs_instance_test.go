package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestClassAndInstance(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-class-vs-instance/problem")

	for _, tt := range []struct {
		age  int
		want string
	}{
		{
			age: -1,
			want: `Age is not valid, setting age to 0.
You are young.
You are young.
`,
		},
		{
			age: 10,
			want: `You are young.
You are a teenager.
`,
		},
		{
			age: 16,
			want: `You are a teenager.
You are old.
`,
		},
		{
			age: 18,
			want: `You are old.
You are old.
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.age), func(t *testing.T) {
			writer := utils.NewStringWriter()
			funcPrintln = func(a ...any) (n int, err error) {
				return fmt.Fprintln(writer, a...)
			}
			defer func() { funcPrintln = fmt.Println }()

			classAndInstance(tt.age)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
