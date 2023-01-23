package p1600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1654(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1654")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 11
802
743
457
539`,
			"200",
		},
		{
			`4 4
200
200
200
1`,
			"100",
		},
		{
			`5 6
1
1
1
1
11`,
			"1",
		},
		{
			`1 4
5`,
			"1",
		},
		{
			`1 2
9`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1654(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
