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

func TestSolve2738(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2738")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 3
1 1 1
2 2 2
0 1 0
3 3 3
4 4 4
5 5 100`,
			`4 4 4
6 6 6
5 6 100
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2700.Solve2738(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
