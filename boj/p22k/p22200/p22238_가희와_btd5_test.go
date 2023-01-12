package p22200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p22k/p22200"
)

func TestSolve22238(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/22238")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 1
1 1 3
3 3 4
2 2 2
1 1 3`,
			`1
`,
		},
		{
			`3 1
0 -1 3
0 -2 4
0 -3 2
0 -3 3`,
			`1
`,
		},
		{
			`3 1
-1 -1 3
-2 -2 4
-3 -3 2
-3 -3 3`,
			`1
`,
		},
		{
			`3 1
1 -1 3
2 -2 4
3 -3 2
3 -3 3`,
			`1
`,
		},
		{
			`3 1
-1 1 3
-2 2 4
-3 3 2
-3 3 3`,
			`1
`,
		},
		{
			`3 1
1 0 3
2 0 4
3 0 2
3 0 3`,
			`1
`,
		},
		{
			`3 1
-1 0 3
-2 0 4
-3 0 2
-3 0 3`,
			`1
`,
		},
		{
			`5 5
1 1 5
2 2 4
3 3 3
4 4 2
5 5 1
1 1 1
2 2 1
3 3 1
4 4 1
5 5 1`,
			`4
3
2
1
0
`,
		},
		{
			`3 3
1 2 6
2 4 12
3 6 18
1 1 100
1 2 10
1 2 20
`,
			`3
2
0
`,
		},
		{
			`3 5
135 2 6
270 4 12
945 14 18
1 1 100
4050 60 5
405 6 1
135 2 5
270 4 1
`,
			`3
3
2
2
1
`,
		},
		{
			`3 3
0 2 6
0 4 12
0 6 18
1 1 100
0 2 10
0 2 20
`,
			`3
2
0
`,
		},
		{
			`3 3
2 0 6
4 0 12
6 0 18
1 1 100
2 0 10
4 0 5
`,
			`3
2
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p22200.Solve22238(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
