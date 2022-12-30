package p1800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1800"
)

func TestSolve1806(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1806")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 15
5 1 3 5 10 7 4 9 2 8`,
			"2",
		},
		{
			`10 11
10 1 1 1 1 1 1 1 1 1`,
			"2",
		},
		{
			`10 10
10 1 1 1 1 1 1 1 1 1`,
			"1",
		},
		{
			`10 100
1 1 1 1 1 1 1 1 1 1`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1800.Solve1806(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
