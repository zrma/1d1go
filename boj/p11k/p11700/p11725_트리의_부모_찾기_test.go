package p11700_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
)

func TestSolve11725(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11725")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1 6
6 3
3 5
4 1
2 4
4 7`,
			`4
6
1
3
1
4
`,
		},
		{
			`12
1 2
1 3
2 4
3 5
3 6
4 7
4 8
5 9
5 10
6 11
6 12`,
			`1
1
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
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p11700.Solve11725(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
