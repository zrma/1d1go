package p2700_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
)

func TestSolve2741(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2741")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`1
`,
		},
		{
			"2",
			`1
2
`,
		},
		{
			"3",
			`1
2
3
`,
		},
		{
			"4",
			`1
2
3
4
`,
		},
		{
			"5",
			`1
2
3
4
5
`,
		},
		{
			"6",
			`1
2
3
4
5
6
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2700.Solve2741(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
