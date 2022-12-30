package p1200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1200"
)

func TestSolve1271(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1271")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1000 100`,
			`10
0
`,
		},
		{
			`1000000000000000000 1000000000000000000`,
			`1
0
`,
		},
		{
			`1000000000000000000000000000000000000000000000000000000 1000000000000000000000000000000000000000000000000000000`,
			`1
0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1200.Solve1271(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
