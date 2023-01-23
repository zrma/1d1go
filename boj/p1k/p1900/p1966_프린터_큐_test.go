package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1966(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1966")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
1 0
5
4 2
1 2 3 4
6 0
1 1 9 1 1 1`,
			`1
2
5
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1966(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
