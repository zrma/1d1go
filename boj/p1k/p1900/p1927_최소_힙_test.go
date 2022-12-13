package p1900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
)

func TestSolve1927(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1927")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`9
0
12345678
1
2
0
0
0
0
32`,
			`0
1
2
12345678
0
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
0
0
0
0
0
0
`,
			`1
2
3
4
5
6
`,
		},
		{
			`12
6
5
4
3
2
1
0
0
0
0
0
0
`,
			`1
2
3
4
5
6
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1900.Solve1927(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
