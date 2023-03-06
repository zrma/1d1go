package p2000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2098(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2098")

	tests := []struct {
		give string
		want string
	}{
		{
			`4
0 10 15 20
5 0 9 10
6 13 0 12
8 8 9 0`,
			"35",
		},
		{
			`4
0 7 3 3
7 0 9 2
1 9 0 12
7 7 12 0`,
			"20",
		},
		{
			`10
0 1 2 3 4 5 6 7 8 0
1 0 1 2 3 4 5 6 7 8
2 1 0 1 2 3 4 5 6 7
3 2 1 0 1 2 3 4 5 6
4 3 2 1 0 1 2 3 4 5
5 4 3 2 1 0 1 2 3 4
6 5 4 3 2 1 0 1 2 3
7 6 5 4 3 2 1 0 1 2
0 7 6 5 4 3 2 1 0 1
9 8 7 6 5 4 3 2 1 0`,
			"18",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2098(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
