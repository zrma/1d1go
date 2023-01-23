package p1800

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test_data/p1865_give.txt
var p1865give string

func TestSolve1865(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1865")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
3 3 1
1 2 2
1 3 4
2 3 1
3 1 3
3 2 1
1 2 3
2 3 4
3 1 8`,
			`NO
YES
`,
		},
		{
			`1
500 1 1
333 444 1
333 444 2`,
			`YES
`,
		},
		{
			`1
4 3 1
1 2 1
2 1 1
3 4 2
4 3 1`,
			`NO
`,
		},
		{
			p1865give,
			`NO
NO
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1865(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
