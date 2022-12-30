package p2800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
)

func TestSolve2845(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2845")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 10
10 10 10 10 10`,
			`0 0 0 0 0 `,
		},
		{
			`5 20
99 101 1000 0 97`,
			`-1 1 900 -100 -3 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2800.Solve2845(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
