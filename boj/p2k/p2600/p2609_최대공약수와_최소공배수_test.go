package p2600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2600"
)

func TestSolve2609(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2609")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"24 18",
			`6
72
`,
		},
		{
			"24 24",
			`24
24
`,
		},
		{
			"12 24",
			`12
24
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2600.Solve2609(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
