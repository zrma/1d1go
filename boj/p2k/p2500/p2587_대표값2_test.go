package p2500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2587(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2587")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
40
30
60
30`,
			`34
30
`,
		},
		{
			`10
20
35
40
50`,
			`31
35
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2587(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
