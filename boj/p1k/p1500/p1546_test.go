package p1500_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1500"
	"1d1go/utils"
)

func TestSolve1546(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1546")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`3
40 80 60`,
			"75.0",
		},
		{
			`3
10 20 30`,
			"66.666667",
		},
		{
			`4
1 100 100 100`,
			"75.25",
		},
		{
			`5
1 2 4 8 16`,
			"38.75",
		},
		{
			`2
3 10`,
			"65.0",
		},
		{
			`4
10 20 0 100`,
			"32.5",
		},
		{
			`1
50`,
			"100.0",
		},
		{
			`9
10 20 30 40 50 60 70 80 90`,
			"55.55555555555556",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1500.Solve1546(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			want, err := strconv.ParseFloat(tt.want, 64)
			assert.NoError(t, err)

			writen := writer.String()
			got, err := strconv.ParseFloat(writen, 64)
			assert.NoError(t, err)

			const epsilon = 1e-6
			assert.InDelta(t, want, got, epsilon)
		})
	}
}
