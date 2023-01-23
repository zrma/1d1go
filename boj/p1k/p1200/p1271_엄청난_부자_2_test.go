package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

			Solve1271(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
