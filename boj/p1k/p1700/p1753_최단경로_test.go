package p1700_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
)

func TestSolve1753(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1753")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 6
1
5 1 1
1 2 2
1 3 3
2 3 4
2 4 5
3 4 6`,
			`0
2
3
7
INF
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1700.Solve1753(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
