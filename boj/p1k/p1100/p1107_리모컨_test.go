package p1100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
)

func TestSolve1107(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1107")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5457
3
6 7 8`,
			"6",
		},
		{
			`100
5
0 1 2 3 4`,
			"0",
		},
		{
			`500000
9
0 1 2 3 5 6 7 8 9`,
			"55562",
		},
		{
			`500000
8
0 2 3 4 6 7 8 9`,
			"11117",
		},
		{
			`500000
10
0 1 2 3 4 5 6 7 8 9`,
			"499900",
		},
		{
			`500000
9
0 1 2 3 4 6 7 8 9`,
			"55561",
		},
		{
			`100
3
1 0 5`,
			"0",
		},
		{
			`14124
0`,
			"5",
		},
		{
			`1
9
1 2 3 4 5 6 7 8 9`,
			"2",
		},
		{
			`80000
2
8 9`,
			"2228",
		},
		{
			`1555
3
0 1 9`,
			"670",
		},
		{
			`140200
3
4 5 6`,
			"207",
		},
		{
			`0
1
0`,
			"2",
		},
		{
			`0
1
1`,
			"1",
		},
		{
			`99
0`,
			"1",
		},
		{
			`102
0`,
			"2",
		},
		{
			`103
0`,
			"3",
		},
		{
			`104
0`,
			"3",
		},
		{
			`105
0`,
			"3",
		},
		{
			`104
1
4`,
			"4",
		},
		{
			`105
1
5`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1100.Solve1107(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
