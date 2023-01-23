package p11700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11780(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11780")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
14
1 2 2
1 3 3
1 4 1
1 5 10
2 4 2
3 4 1
3 5 1
4 5 3
3 5 10
3 1 8
1 4 2
5 1 7
3 4 2
5 2 4`,
			`0 2 3 1 4
12 0 15 2 5
8 5 0 1 1
10 7 13 0 3
7 4 10 6 0
0
2 1 2
2 1 3
2 1 4
3 1 3 5
4 2 4 5 1
0
5 2 4 5 1 3
2 2 4
3 2 4 5
2 3 1
3 3 5 2
0
2 3 4
2 3 5
3 4 5 1
3 4 5 2
4 4 5 1 3
0
2 4 5
2 5 1
2 5 2
3 5 1 3
3 5 2 4
0
`,
		},
		{
			`5
1
1 2 3`,
			`0 3 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0
2 1 2
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11780(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
