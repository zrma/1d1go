package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1202(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1202")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2 1
5 10
100 100
11`,
			"10",
		},
		{
			`3 2
1 65
5 23
2 99
10
2`,
			"164",
		},
		{
			`2 3
1 1
2 2
2 2
3 3
4 4`,
			"3",
		},
		{
			`3 2
10 65
5 23
2 99
10
2`,
			"164",
		},
		{
			`3 2
10 65
5 99
2 23
10
2`,
			"122",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1202(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
