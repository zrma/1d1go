package p1600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1600"
	"1d1go/utils"
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
			writer := utils.NewStringWriter()

			p1600.Solve1654(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
