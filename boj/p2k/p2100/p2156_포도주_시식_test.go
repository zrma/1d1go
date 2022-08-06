package p2100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2100"
	"1d1go/utils"
)

func TestSolve2156(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2156")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
6
10
13
9
8
1`,
			"33",
		},
		{
			`7
6
10
13
9
8
1
2`,
			"35",
		},
		{
			`8
6
10
13
9
8
1
2
101`,
			"136",
		},
		{
			`1
123`,
			"123",
		},
		{
			`2
12
34`,
			"46",
		},
		{
			`3
1
2
3`,
			"5",
		},
		{
			`4
1
2
3
4`,
			"8",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2100.Solve2156(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
