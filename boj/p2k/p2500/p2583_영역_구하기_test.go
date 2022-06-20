package p2500_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2583(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2583")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 7 3
0 2 4 4
1 1 2 5
4 0 6 2`,
			`3
1 7 13 `,
		},
		{
			`100 100 0`,
			`1
10000 `,
		},
		{
			`100 100 1
0 0 1 1`,
			`1
9999 `,
		},
		{
			`100 100 1
1 0 2 100 `,
			`2
100 9800 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p2500.Solve2583(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
