package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1976(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1976")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
3
0 1 0
1 0 1
0 1 0
1 2 3`,
			"YES",
		},
		{
			`5
3
0 1 0 0 0
1 0 1 0 0
0 1 0 0 0
0 0 0 0 0
0 0 0 0 0
1 2 3
`,
			"YES",
		},
		{
			`5
4
0 1 0 0 0
1 0 1 0 0
0 1 0 0 0
0 0 0 0 0
0 0 0 0 0
1 2 3 4
`,
			"NO",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d/UnionFind", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1976(reader, writer, Solve1976UnionFind)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/DFS", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1976(reader, writer, Solve1976DFS)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/BFS", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1976(reader, writer, Solve1976BFS)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
