package p2400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2437(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2437")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
3 1 6 2 7 30 1`,
			"21",
		},
		{
			`7
1 1 2 3 6 7 30`,
			"21",
		},
		{
			`1
1`,
			"2",
		},
		{
			`1
2`,
			"1",
		},
		{
			`1
3`,
			"1",
		},
		{
			`2
1 1`,
			"3",
		},
		{
			`2
1 2`,
			"4",
		},
		{
			`2
1 3`,
			"2",
		},
		{
			`2
1 4`,
			"2",
		},
		{
			`3
1 2 4`,
			"8",
		},
		{
			`3
1 2 5`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2437(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
