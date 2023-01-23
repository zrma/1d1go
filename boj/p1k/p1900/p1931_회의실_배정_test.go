package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1931(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1931")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`11
1 4
3 5
0 6
5 7
3 8
5 9
6 10
8 11
8 12
2 13
12 14`,
			"4",
		},
		{
			`2
1 2
1 2`,
			"1",
		},
		{
			`4
1 2
1 2
1 2
1 2`,
			"1",
		},
		{
			`4
1 2
0 0
1 1
1 1`,
			"4",
		},
		{
			`6
0 0
0 3
0 1
1 1
1 2
2 3`,
			"5",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1931(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
