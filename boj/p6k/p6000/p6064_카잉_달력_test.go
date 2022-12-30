package p6000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p6k/p6000"
)

func TestSolve6064(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/6064")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
10 12 3 9
10 12 7 2
13 11 5 6`,
			`33
-1
83
`,
		},
		{
			`15
3 5 1 1
3 5 2 2
3 5 3 3
3 5 1 4
3 5 2 5
3 5 3 1
3 5 1 2
3 5 2 3
3 5 3 4
3 5 1 5
3 5 2 1
3 5 3 2
3 5 1 3
3 5 2 4
3 5 3 5
`,
			`1
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
12
13
14
15
`,
		},
		{
			`15
5 3 1 1
5 3 2 2
5 3 3 3
5 3 4 1
5 3 5 2
5 3 1 3
5 3 2 1
5 3 3 2
5 3 4 3
5 3 5 1
5 3 1 2
5 3 2 3
5 3 3 1
5 3 4 2
5 3 5 3
`,
			`1
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
12
13
14
15
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p6000.Solve6064(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
