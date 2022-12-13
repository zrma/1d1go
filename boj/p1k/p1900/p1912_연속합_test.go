package p1900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
)

func TestSolve1912(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1912")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
10 -4 3 1 5 6 -35 12 21 -1`,
			"33",
		},
		{
			`10
2 1 -4 3 4 -4 6 5 -5 1`,
			"14",
		},
		{
			`5
-1 -2 -3 -4 -5`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1900.Solve1912(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
