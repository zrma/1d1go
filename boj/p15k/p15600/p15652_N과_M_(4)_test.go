package p15600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15600"
)

func TestSolve15652(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15652")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"3 1",
			`1
2
3
`,
		},
		{
			"4 2",
			`1 1
1 2
1 3
1 4
2 2
2 3
2 4
3 3
3 4
4 4
`,
		},
		{
			"3 3",
			`1 1 1
1 1 2
1 1 3
1 2 2
1 2 3
1 3 3
2 2 2
2 2 3
2 3 3
3 3 3
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p15600.Solve15652(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
