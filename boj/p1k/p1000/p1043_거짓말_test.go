package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1043(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1043")

	tests := []struct {
		give string
		want string
	}{
		{
			`4 3
0
2 1 2
1 3
3 2 3 4`,
			"3",
		},
		{
			`4 1
1 1
4 1 2 3 4`,
			"0",
		},
		{
			`4 5
1 1
1 1
1 2
1 3
1 4
2 4 1`,
			"2",
		},
		{
			`10 9
4 1 2 3 4
2 1 5
2 2 6
1 7
1 8
2 7 8
1 9
1 10
2 3 10
1 4`,
			"4",
		},
		{
			`8 5
3 1 2 7
2 3 4
1 5
2 5 6
2 6 8
1 8`,
			"5",
		},
		{
			`3 4
1 3
1 1
1 2
2 1 2
3 1 2 3`,
			"0",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1043(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
