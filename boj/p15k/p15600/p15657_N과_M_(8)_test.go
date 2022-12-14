package p15600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15600"
)

func TestSolve15657(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15657")

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
			`1 1
1 7
1 8
1 9
7 7
7 8
7 9
8 8
8 9
9 9
`,
		},
		{
			`4 4
1231 1232 1233 1234`,
			`1231 1231 1231 1231
1231 1231 1231 1232
1231 1231 1231 1233
1231 1231 1231 1234
1231 1231 1232 1232
1231 1231 1232 1233
1231 1231 1232 1234
1231 1231 1233 1233
1231 1231 1233 1234
1231 1231 1234 1234
1231 1232 1232 1232
1231 1232 1232 1233
1231 1232 1232 1234
1231 1232 1233 1233
1231 1232 1233 1234
1231 1232 1234 1234
1231 1233 1233 1233
1231 1233 1233 1234
1231 1233 1234 1234
1231 1234 1234 1234
1232 1232 1232 1232
1232 1232 1232 1233
1232 1232 1232 1234
1232 1232 1233 1233
1232 1232 1233 1234
1232 1232 1234 1234
1232 1233 1233 1233
1232 1233 1233 1234
1232 1233 1234 1234
1232 1234 1234 1234
1233 1233 1233 1233
1233 1233 1233 1234
1233 1233 1234 1234
1233 1234 1234 1234
1234 1234 1234 1234
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p15600.Solve15657(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
