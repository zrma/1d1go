package p1900_test

import (
	"fmt"
	"testing"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve1912(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1912")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
10 -4 3 1 5 6 -35 12 21 -1`,
			"33",
		},
		{
			`10
2 1 -4 3 4 -4 6 5 -5 1`,
			"14",
		},
		{
			`5
-1 -2 -3 -4 -5`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p1900.Solve1912(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
