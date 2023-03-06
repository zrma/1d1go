package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1780(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1780")

	tests := []struct {
		give string
		want string
	}{
		{
			`9
0 0 0 1 1 1 -1 -1 -1
0 0 0 1 1 1 -1 -1 -1
0 0 0 1 1 1 -1 -1 -1
1 1 1 0 0 0 0 0 0
1 1 1 0 0 0 0 0 0
1 1 1 0 0 0 0 0 0
0 1 -1 0 1 -1 0 1 -1
0 -1 1 0 1 -1 0 1 -1
0 1 -1 1 0 -1 0 1 -1`,
			`10
12
11
`,
		},
		{
			`3
-1 0 1
-1 0 1
-1 0 1`,
			`3
3
3
`,
		},
		{
			`3
1 1 1
1 1 1
1 1 1`,
			`0
0
1
`,
		},
		{
			`9
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1`,
			`0
0
1
`,
		},
		{
			`3
0 0 0
0 0 0
0 0 1`,
			`0
8
1
`,
		},
		{
			`9
1 1 1 1 1 1 1 1 1
1 1 1 0 0 0 1 1 1
1 1 1 0 0 0 1 1 1
-1 -1 -1 -1 -1 -1 -1 -1 -1
-1 -1 -1 -1 -1 -1 -1 -1 -1
-1 -1 -1 -1 -1 -1 -1 -1 -1
0 0 0 1 1 1 1 1 1
0 0 0 0 0 0 1 1 1
0 0 0 -1 -1 -1 1 1 1`,
			`6
10
9
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1780(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
