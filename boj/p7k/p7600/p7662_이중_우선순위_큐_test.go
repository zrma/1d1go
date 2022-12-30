package p7600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p7k/p7600"
)

func TestSolve7662(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7662")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
7
I 16
I -5643
D -1
D 1
D 1
I 123
D -1
9
I -45
I 653
D 1
I -642
I 45
I 97
D 1
D -1
I 333`,
			`EMPTY
333 -45
`,
		},
		{
			`1
9
I 1
I 2
I 3
I 4
D 1
D 1
D 1
D 1
I 2`,
			`2 2
`,
		},
		{
			`1
7
I 1
I 1
I 1
D 1
D 1
D 1
D 1`,
			`EMPTY
`,
		},
		{
			`1
7
I 1
I 1
I 1
D -1
D -1
D -1
D -1`,
			`EMPTY
`,
		},
		{
			`1
11
I 1
I 1
I 1
I 1
I 1
I 1
D 1
D 1
D -1
D -1
D 1
`,
			`1 1
`,
		},
		{
			`1
1
D 1`,
			`EMPTY
`,
		},
		{
			`1
1
D -1`,
			`EMPTY
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p7600.Solve7662(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
