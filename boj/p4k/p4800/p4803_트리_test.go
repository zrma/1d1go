package p4800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4800"
)

func TestSolve4803(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4803")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6 3
1 2
2 3
3 4
6 5
1 2
2 3
3 4
4 5
5 6
6 6
1 2
2 3
1 3
4 5
5 6
6 4
0 0`,
			`Case 1: A forest of 3 trees.
Case 2: There is one tree.
Case 3: No trees.
`,
		},
		{
			`6 7
1 2
1 3
1 4
2 3
2 4
3 4
5 6
0 0`,
			`Case 1: There is one tree.
`,
		},
		{
			`3 3
1 1
2 3
1 2
0 0`,
			`Case 1: No trees.
`,
		},
	} {
		t.Run(fmt.Sprintf("%d/UnionFind", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p4800.Solve4803(reader, writer, p4800.Solve4803UnionFind)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/DFS", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p4800.Solve4803(reader, writer, p4800.Solve4803DFS)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/BFS", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p4800.Solve4803(reader, writer, p4800.Solve4803BFS)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
