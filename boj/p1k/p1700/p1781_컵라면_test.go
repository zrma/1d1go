package p1700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
)

func TestSolve1781(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1781")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1 6
1 7
3 2
3 1
2 4
2 5
6 1`,
			"15",
		},
		{
			`3
3 3
3 2
3 1`,
			"6",
		},
		{
			`4
3 3
3 2
3 4
2 4`,
			"11",
		},
		{
			`5
3 1
3 2
2 3
4 4
4 5`,
			"14",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1700.Solve1781(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
