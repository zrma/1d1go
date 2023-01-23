package p1600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1655(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1655")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1
5
2
10
-99
7
5`,
			`1
1
2
2
2
2
5
`,
		},
		{
			`10
9
5
8
4
7
3
6
2
5
1`,
			`9
5
8
5
7
5
6
5
5
5
`,
		},
		{
			`12
1
2
3
4
5
6
7
8
9
10
11
12`,
			`1
1
2
2
3
3
4
4
5
5
6
6
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1655(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
