package p1000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
)

func TestSolve1015(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1015")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
2 3 1`,
			"1 2 0",
		},
		{
			`4
2 1 3 1`,
			"2 0 3 1",
		},
		{
			`8
4 1 6 1 3 6 1 4`,
			"4 0 6 1 3 7 2 5",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1000.Solve1015(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
