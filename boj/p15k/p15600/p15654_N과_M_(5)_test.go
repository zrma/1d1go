package p15600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve15654(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15654")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 1
4 5 2`,
			`2
4
5
`,
		},
		{
			`4 2
9 8 7 1`,
			`1 7
1 8
1 9
7 1
7 8
7 9
8 1
8 7
8 9
9 1
9 7
9 8
`,
		},
		{
			`4 4
1231 1232 1233 1234`,
			`1231 1232 1233 1234
1231 1232 1234 1233
1231 1233 1232 1234
1231 1233 1234 1232
1231 1234 1232 1233
1231 1234 1233 1232
1232 1231 1233 1234
1232 1231 1234 1233
1232 1233 1231 1234
1232 1233 1234 1231
1232 1234 1231 1233
1232 1234 1233 1231
1233 1231 1232 1234
1233 1231 1234 1232
1233 1232 1231 1234
1233 1232 1234 1231
1233 1234 1231 1232
1233 1234 1232 1231
1234 1231 1232 1233
1234 1231 1233 1232
1234 1232 1231 1233
1234 1232 1233 1231
1234 1233 1231 1232
1234 1233 1232 1231
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve15654(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
