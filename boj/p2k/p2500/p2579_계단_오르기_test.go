package p2500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2579(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
