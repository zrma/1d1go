package p1800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1800"
)

func TestSolve1802(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1802")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
0
000
1000110`,
			`YES
NO
YES
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1800.Solve1802(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
