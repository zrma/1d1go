package p2100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2100"
)

func TestSolve2108(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2108")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
1
3
8
-2
2`,
			`2
2
1
10
`,
		},
		{
			`1
4000`,
			`4000
4000
4000
0
`,
		},
		{
			`5
-1
-2
-3
-1
-2`,
			`-2
-2
-1
2
`,
		},
		{
			`3
0
0
-1`,
			`0
0
0
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2100.Solve2108(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
