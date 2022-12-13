package p2500_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
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
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2500.Solve2587(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
