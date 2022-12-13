package p1100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
)

func TestSolve1167(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1167")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
1 3 2 -1
2 4 4 -1
3 1 2 4 3 -1
4 2 4 3 3 5 6 -1
5 4 6 -1`,
			"11",
		},
		{
			`12
1 2 3 3 2 -1
2 1 3 4 5 -1
3 1 2 5 11 6 9 -1
4 2 5 7 1 8 7 -1
5 3 11 9 15 10 4 -1
6 3 9 11 6 12 10 -1
7 4 1 -1
8 4 7 -1
9 5 15 -1
10 5 4 -1
11 6 6 -1
12 6 10 -1`,
			"45",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1100.Solve1167(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
