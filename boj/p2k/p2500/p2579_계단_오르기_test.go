package p2500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2579(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2579")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
10
20
15
25
10
20`,
			"75",
		},
		{
			`1
10`,
			"10",
		},
		{
			`2
5
10`,
			"15",
		},
		{
			`3
10
20
30`,
			"50",
		},
		{
			`4
10
20
30
40`,
			"80",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2500.Solve2579(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
