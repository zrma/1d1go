package p2600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2600"
)

func TestSolve2667(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2667")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
0110100
0110101
1110101
0000111
0100000
0111110
0111000`,
			`3
7
8
9
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2600.Solve2667(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
